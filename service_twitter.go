package main

import (
	"log"
	"unicode/utf8"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func tweet(c *Content) int64 {
	log.Println("Tweeting content")
	client := getOriginClient()
	verify(client)
	tweet_id := tweetContent(client, c)

	client2 := getRepeatClient()
	reTweetContent(client2, tweet_id)
	return tweet_id
}

func getOriginClient() *twitter.Client {
	return getConfiguredTwitterClient(
		env(TWITTER_ORIGIN_CONSUMER_KEY),
		env(TWITTER_ORIGIN_CONSUMER_SECRET),
		env(TWITTER_ORIGIN_ACCESS_TOKEN),
		env(TWITTER_ORIGIN_ACCESS_TOKEN_SECRET))
}

func getRepeatClient() *twitter.Client {
	return getConfiguredTwitterClient(
		env(TWITTER_REPEAT_CONSUMER_KEY),
		env(TWITTER_REPEAT_CONSUMER_SECRET),
		env(TWITTER_REPEAT_ACCESS_TOKEN),
		env(TWITTER_REPEAT_ACCESS_TOKEN_SECRET))
}

func getConfiguredTwitterClient(
	consumerKey,
	consumerSecret,
	accessToken,
	accessTokenSecret string) *twitter.Client {

	cred := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	client := cred.Client(oauth1.NoContext, token)
	return twitter.NewClient(client)
}

func verify(client *twitter.Client) {
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	log.Printf("User's ACCOUNT:\n%+v\n", user)
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
