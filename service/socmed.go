package service

import (
	"fmt"
	"log"
	"unicode/utf8"

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

	log.Printf("Creating socmed service at localhost:8080 -- access with http://localhost:8080%s\n", path)

	service.Route(service.POST("/twitter").Filter(basicAuthenticate).To(Tweet))
	service.Route(service.POST("/tumblr").Filter(basicAuthenticate).To(PostToTumblr))

	return service
}

func Tweet(request *restful.Request, response *restful.Response) {
	cred := oauth1.NewConfig(
		config.GetTwitterConsumerKey(),
		config.GetTwitterConsumerSecret())

	token := oauth1.NewToken(
		config.GetTwitterAccessToken(),
		config.GetTwitterAccessTokenSecret())

	httpClient := cred.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

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
	tweet, _, _ := client.Statuses.Update(getTweetText(), nil)
	fmt.Printf("Posted tweet \n%v\n", tweet)
}

func getTweetText() string {
	url := "https://devabo.de"
	tweet := "replace me" + url

	for _, tag := range config.GetTags() {
		if utf8.RuneCountInString(tweet+" "+tag) > 140 {
			return tweet
		}
		tweet += " " + tag
	}

	return tweet
}

func PostToTumblr(request *restful.Request, response *restful.Response) {
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
	tags := "comic,webcomic,graphicnovel,drawing,art,narrative,scifi,sci-fi,science-fiction,dystopy,parody,humor,nerd,pulp,geek,blackandwhite"
	prodUrl := "https://devabo.de ..."

	photoPostByURL := client.CreatePhoto(
		blogname,
		map[string]string{
			"link":    prodUrl,
			"source":  "replace me",
			"caption": "replace me",
			"tags":    tags,
			"state":   state})
	if photoPostByURL == nil {
		fmt.Println("done")
	} else {
		fmt.Println(photoPostByURL)
	}
}
