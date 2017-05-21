package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicSocMed/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
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

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	fbResponse, _ := httpClient.Head(url)
	for k, v := range fbResponse.Header {
		fmt.Printf("key:%s, value:%s\n", k, v)
	}

	//rbody, err := ioutil.ReadAll(fbResponse.Body)
	//http.Redirect(response, request.Request, url, http.StatusTemporaryRedirect)
}

//func Publish
/*
func FacebookInit(w http.ResponseWriter, r *http.Request) {
	Url, err := url.Parse(oauthConf.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", oauthConf.ClientID)
	parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
	parameters.Add("redirect_uri", oauthConf.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", oauthStateString)
	Url.RawQuery = parameters.Encode()
	url := Url.String()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
*/

func FacebookCallback(r *restful.Request, response *restful.Response) {

	state := r.Request.FormValue("state")
	code := r.Request.FormValue("code")

	fmt.Println(state)
	fmt.Println(code)

	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		return
	}

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		return
	}
	defer resp.Body.Close()

	response2, _ := ioutil.ReadAll(resp.Body)

	log.Printf("parseResponseBody: %s\n", string(response2))
}

func logRequest(r *restful.Request) {

	// Create return string
	var req []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Request.Method, r.Request.URL, r.Request.Proto)
	req = append(req, url)
	// Add the host
	req = append(req, fmt.Sprintf("Host: %v", r.Request.Host))
	// Loop through headers
	for name, headers := range r.Request.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			req = append(req, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Request.Method == "POST" {
		r.Request.ParseForm()
		req = append(req, "\n")
		req = append(req, r.Request.Form.Encode())
	}
	// Return the request as a string
	fmt.Println(strings.Join(req, "\n"))
}

/*
func FacebookCallback(w http.ResponseWriter, r *http.Request) {

}
*/
