package service

import (
	"log"
	"unicode/utf8"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ingmardrewing/gomicSocMed/config"
)

func tweet(c *Content) {
	log.Println("Tweeting content")
	client := getTwitterClient()
	verify(client)
	tweetContent(client, c)
}

func verify(client *twitter.Client) {
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	log.Printf("User's ACCOUNT:\n%+v\n", user)
}

func getTwitterClient() *twitter.Client {
	cred := oauth1.NewConfig(
		config.GetTwitterConsumerKey(),
		config.GetTwitterConsumerSecret())
	token := oauth1.NewToken(
		config.GetTwitterAccessToken(),
		config.GetTwitterAccessTokenSecret())

	httpClient := cred.Client(oauth1.NoContext, token)
	return twitter.NewClient(httpClient)
}

func tweetContent(client *twitter.Client, c *Content) {
	if config.IsProd() {
		tweet, _, _ := client.Statuses.Update(getTweetText(c), nil)
		log.Printf("Posted tweet \n%v\n", tweet)
	}
}

func getTweetText(c *Content) string {
	tweet := c.Title + " " + c.Link

	for _, tag := range c.Tags {
		if utf8.RuneCountInString(tweet+" #"+tag) > 140 {
			return tweet
		}
		tweet += " #" + tag
	}

	return tweet
}
