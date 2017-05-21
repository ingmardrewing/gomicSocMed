package service

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	restful "github.com/emicklei/go-restful"
	fb "github.com/huandu/facebook"
	"github.com/ingmardrewing/gomicSocMed/config"
	"github.com/ingmardrewing/gomicSocMed/db"
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
		ClientID:     config.GetFacebookApplicationId(),
		ClientSecret: config.GetFacebookApplicationSecret(),
		RedirectURL:  config.GetFacebookCallbackUrl(),
		Scopes:       []string{"public_profile"},
		Endpoint:     facebook.Endpoint,
	}
	src = rand.NewSource(time.Now().UnixNano())
)

func getRandomString(n int) string {
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

func FacebookInit(request *restful.Request, response *restful.Response) {
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

	http.Redirect(response, request.Request, url, http.StatusTemporaryRedirect)
}

func FacebookCallback(r *restful.Request, response *restful.Response) {
	code := r.Request.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		return
	}

	accounts_url := "https://graph.facebook.com/me/accounts?access_token=" + url.QueryEscape(token.AccessToken)

	resp, err := http.Get(accounts_url)
	if err != nil {
		log.Printf("Get: %s\n", err)
		return
	}
	defer resp.Body.Close()

	response2, _ := ioutil.ReadAll(resp.Body)

	storeFBAccessToken(string(response2))
}

func postToFacebook(c *Content) {
	log.Println("Posting to facebook")

	access_token := retrieveFBAccessTokens()
	page_id := config.GetFacebookPageId()

	config.GetFacebookPageId()
	_, err := fb.Post("/"+page_id+"/feed", fb.Params{
		"type":         "link",
		"name":         c.Title,
		"caption":      c.Title,
		"picture":      c.ImgUrl,
		"link":         c.Link,
		"description":  c.Description,
		"message":      getTagsForFacebook(c),
		"access_token": access_token,
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Facebook posting succeeded")
	}
}

func getTagsForFacebook(c *Content) string {
	txt := ""

	for _, tag := range c.Tags {
		if utf8.RuneCountInString(txt+" #"+tag) > 140 {
			return txt
		}
		txt += " #" + tag
	}

	return txt
}

func storeFBAccessToken(token string) {
	// TODO parse json properly, store all tokens
	re := regexp.MustCompile("\"access_token\":\"([^\"]+)\"")
	matches := re.FindStringSubmatch(token)
	if db.TokenExists("fb_devabode") {
		db.UpdateToken("fb_devabode", matches[1])
	} else {
		db.InsertToken("fb_devabode", matches[1])
	}
}

func retrieveFBAccessTokens() string {
	return db.GetToken("fb_devabode")
}
