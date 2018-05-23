package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	restful "github.com/emicklei/go-restful"
	fb "github.com/huandu/facebook"
	store "github.com/ingmardrewing/fsKeyValueStore"
	shared "github.com/ingmardrewing/gomicSocMedShared"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     shared.Env(shared.FB_APPLICATION_ID),
		ClientSecret: shared.Env(shared.FB_APPLICATION_SECRET),
		RedirectURL:  shared.Env(shared.FB_CALLBACK_URL),
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	src = rand.NewSource(time.Now().UnixNano())
)

func getRandomString(n int) string {
	log.Println("getRandomString")
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func FacebookGetAccessToken(request *restful.Request, response *restful.Response) {
	log.Println("FacebookGetAccessToken")
	Url, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}

	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", getRandomString(18))

	Url.RawQuery = parameters.Encode()
	url := Url.String()

	log.Println("Redirecting to", url)
	http.Redirect(response, request.Request, url, http.StatusTemporaryRedirect)
}

func FacebookCallback(r *restful.Request, response *restful.Response) {
	log.Println("FacebookCallback called")

	code := r.Request.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		return
	}

	accounts_url := "https://graph.facebook.com/me/accounts?access_token=" + url.QueryEscape(token.AccessToken)
	log.Println("Getting url", accounts_url)

	resp, err := http.Get(accounts_url)
	if err != nil {
		log.Printf("Get: %s\n", err)
		return
	}
	defer resp.Body.Close()

	response2, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(response2))

	extractedToken := extractFBAccessToken(string(response2))
	accountName := "drewingde"
	storeFBAccessToken(accountName, extractedToken)
}

func postToFacebook(c *Content) []fb.Result {
	log.Println("Posting to facebook")
	resp := postToFacebookAsMe(c, "ingmardrewing")
	return []fb.Result{resp}
}

/*
func postToFacebookCascade(c *Content) []fb.Result {
	log.Println("Posting to facebook, cascading")

	results := []fb.Result{}

	resp1 := postToFacebookAs(c, "devabode")
	results = append(results, resp1)

	str, ok := resp1.Get("id").(string)
	if ok {
		parts := strings.Split(str, "_")
		link := fmt.Sprintf("https://www.facebook.com/%s/posts/%s", parts[0], parts[1])

		resp2 := repostToFacebookAs(link, "drewingde")
		results = append(results, resp2)
	}

	return results
}

*/
func postToFacebookAsMe(c *Content, name string) fb.Result {
	log.Println("postToFacebook")
	access_token := retrieveTokenFor(name)
	log.Println("got fb access token", access_token)

	message := c.Title + " " + getTagsForFacebook(c)
	resp, err := fb.Post("/me/feed", fb.Params{
		"type":         "link",
		"name":         c.Title,
		"caption":      c.Title,
		"picture":      c.ImgUrl,
		"link":         c.Link,
		"description":  c.Description,
		"message":      message,
		"access_token": access_token,
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
	return resp
}

/*
func postToFacebookAs(c *Content, name string) fb.Result {
	log.Println("postToFacebook")

	access_token := retrieveTokenFor(name)
	fbpath := "/" + retrieveIdFor(name) + "/feed"
	message := c.Description + " " + getTagsForFacebook(c)

	fmt.Println("posting to " + fbpath)
	resp, err := fb.Post(fbpath, fb.Params{
		"type":         "link",
		"name":         c.Title,
		"caption":      c.Title,
		"picture":      c.ImgUrl,
		"link":         c.Link,
		"description":  c.Description,
		"message":      message,
		"access_token": access_token,
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
	return resp
}
*/

func repostToFacebookAs(link string, name string) fb.Result {
	log.Println("repostToFacebookAs")
	access_token := retrieveTokenFor(name)
	id := retrieveIdFor(name)

	resp, err := fb.Post("/"+id+"/feed", fb.Params{
		"type":         "link",
		"link":         link,
		"access_token": access_token,
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(resp)
	}
	return resp
}

func getTagsForFacebook(c *Content) string {
	log.Println("getTagsForFacebook")
	txt := ""

	for _, tag := range c.Tags {
		txt += " #" + tag
	}

	return txt
}

func extractFBAccessToken(response string) string {
	re := regexp.MustCompile("\"access_token\":\"([^\"]+)\"")
	matches := re.FindStringSubmatch(response)
	return matches[1]
}

func storeFBAccessToken(account_name, token string) {
	fileKey := account_name + "_fb_token"
	store.CreateIfNonExistentElseUpdate(fileKey, token)
}

func retrieveTokenFor(account_name string) string {
	return read(account_name + "_fb_token")
}

func retrieveIdFor(account_name string) string {
	return read(account_name + "_fb_id")
}

func read(key string) string {
	value, err := store.Read(key)
	if err != nil {
		log.Fatal("Couldn't read value for key", key)
	}
	return value
}
