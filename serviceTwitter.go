package main

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	shared "github.com/ingmardrewing/gomicSocMedShared"
)

func tweet(c *shared.Content) int64 {
	log.Println("Tweeting content")
	client := getRepeatClient()
	verify(client)
	return tweetContent(client, c)
}

func doubletweet(c *shared.Content) int64 {
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
		shared.Env(shared.TWITTER_ORIGIN_CONSUMER_KEY),
		shared.Env(shared.TWITTER_ORIGIN_CONSUMER_SECRET),
		shared.Env(shared.TWITTER_ORIGIN_ACCESS_TOKEN),
		shared.Env(shared.TWITTER_ORIGIN_ACCESS_TOKEN_SECRET))
}

func getRepeatClient() *twitter.Client {
	return getConfiguredTwitterClient(
		shared.Env(shared.TWITTER_REPEAT_CONSUMER_KEY),
		shared.Env(shared.TWITTER_REPEAT_CONSUMER_SECRET),
		shared.Env(shared.TWITTER_REPEAT_ACCESS_TOKEN),
		shared.Env(shared.TWITTER_REPEAT_ACCESS_TOKEN_SECRET))
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

func tweetContent(client *twitter.Client, c *shared.Content) int64 {
	tweet, _, _ := client.Statuses.Update(getTweetText(c), nil)
	return tweet.ID
}

func reTweetContent(client *twitter.Client, tweet_id int64) int64 {
	tweet, _, _ := client.Statuses.Retweet(tweet_id, nil)
	return tweet.ID
}

func getTweetText(c *shared.Content) string {
	tweet := c.Title + " " + c.Link
	return addTags(tweet, c.Tags)
}

func addTags(tweet string, tags []string) string {
	TWEET_LIMIT := 280
	for _, tag := range tags {
		tweetPlusTag := tweet + " #" + tag
		if len(tweetPlusTag) > TWEET_LIMIT {
			return tweet
		}
		tweet = tweetPlusTag
	}
	return tweet
}
