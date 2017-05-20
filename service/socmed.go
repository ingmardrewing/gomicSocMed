package service

import (
	"fmt"
	"log"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"

	"github.com/MariaTerzieva/gotumblr"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicSocMed/config"
)

func NewSocMedService() *restful.WebService {
	path := "/0.1/gomic/socmed"
	service := new(restful.WebService)
	service.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Rest base path: %s\n", path)

	service.Route(service.POST("/publish").Filter(basicAuthenticate).To(Publish))

	return service
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	err := authenticate(req)
	if err != nil {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(req, resp)
}

func authenticate(req *restful.Request) error {
	user, pass, _ := req.Request.BasicAuth()
	given_pass := []byte(pass)
	stored_hash := []byte(config.GetPasswordHashForUser(user))
	return bcrypt.CompareHashAndPassword(stored_hash, given_pass)
}

type Content struct {
	Link, ImgUrl, Title, Tags string
}

func Publish(request *restful.Request, response *restful.Response) {
	p := new(Content)
	request.ReadEntity(p)
	response.WriteEntity(p)
}

func tweet(request *restful.Request, response *restful.Response) {
	cred := oauth1.NewConfig(
		config.GetTwitterConsumerKey(),
		config.GetTwitterConsumerSecret())

	token := oauth1.NewToken(
		config.GetTwitterAccessToken(),
		config.GetTwitterAccessTokenSecret())

	httpClient := cred.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	if client != nil {
	}

	/*
		verifyParams := &twitter.AccountVerifyParams{
			SkipStatus:   twitter.Bool(true),
			IncludeEmail: twitter.Bool(true),
		}
		user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
		fmt.Printf("User's ACCOUNT:\n%+v\n", user)

		fmt.Println(getTweetText())
	*/

	// actually tweet
	// tweet, _, _ := client.Statuses.Update(getTweetText(), nil)
	// fmt.Printf("Posted tweet \n%v\n", tweet)
}

func getTweetText(tags []string) string {
	url := "https://devabo.de"
	tweet := "replace me" + url

	for _, tag := range tags {
		if utf8.RuneCountInString(tweet+" "+tag) > 140 {
			return tweet
		}
		tweet += " " + tag
	}

	return tweet
}

func postToTumblr(request *restful.Request, response *restful.Response) {
	fmt.Println("Post to tumblr")

	cons_key := config.GetTumblrConsumerKey()
	cons_secret := config.GetTumblrConsumerSecret()
	token := config.GetTumblrToken()
	token_secret := config.GetTumblrTokenSecret()

	tumblr_callback_url := "http://localhost/~drewing/cgi-bin/tumblr.pl"

	client := gotumblr.NewTumblrRestClient(
		cons_key,
		cons_secret,
		token,
		token_secret,
		tumblr_callback_url,
		"http://api.tumblr.com")

	blogname := "devabo-de.tumblr.com"
	state := "published"
	//tags := "comic,webcomic,graphicnovel,drawing,art,narrative,scifi,sci-fi,science-fiction,dystopy,parody,humor,nerd,pulp,geek,blackandwhite"
	prodUrl := "https://devabo.de ..."

	if client != nil && len(blogname) > 0 && len(state) > 0 && len(prodUrl) > 0 {
	}

	/*
		photoPostByURL := client.CreatePhoto(
			blogname,
			map[string]string{
				"link":    prodUrl,
				"source":  "imgUrl",
				"caption": "title",
				"tags":    tags,
				"state":   state})
		if photoPostByURL == nil {
			fmt.Println("done")
		} else {
			fmt.Println(photoPostByURL)
		}
	*/
}
