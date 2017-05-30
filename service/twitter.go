package service

import (
	"log"
	"unicode/utf8"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ingmardrewing/gomicSocMed/config"
)

func tweet(c *Content) int64 {
	log.Println("Tweeting content")
	client := getTwitterClient("twitter_devabode")
	verify(client)
	tweet_id := tweetContent(client, c)

	client2 := getTwitterClient("twitter_ingmardrewing")
	reTweetContent(client2, tweet_id)
	return tweet_id
}

func verify(client *twitter.Client) {
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	log.Printf("User's ACCOUNT:\n%+v\n", user)
}

func getTwitterClient(client_name string) *twitter.Client {
	if client_name == "twitter_devabode" {
		cred := oauth1.NewConfig(
			config.GetTwitterConsumerKey(),
			config.GetTwitterConsumerSecret())

		token := oauth1.NewToken(
			config.GetTwitterAccessToken(),
			config.GetTwitterAccessTokenSecret())

		client := cred.Client(oauth1.NoContext, token)
		return twitter.NewClient(client)
	}

	cred := oauth1.NewConfig(
		config.GetTwitterRepeatConsumerKey(),
		config.GetTwitterRepeatConsumerSecret())

	token := oauth1.NewToken(
		config.GetTwitterRepeatAccessToken(),
		config.GetTwitterRepeatAccessTokenSecret())

	client := cred.Client(oauth1.NoContext, token)
	return twitter.NewClient(client)
}

func tweetContent(client *twitter.Client, c *Content) int64 {
	tweet, _, _ := client.Statuses.Update(getTweetText(c), nil)
	return tweet.ID
}

func reTweetContent(client *twitter.Client, tweet_id int64) int64 {
	tweet, _, _ := client.Statuses.Retweet(tweet_id, nil)
	return tweet.ID
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
