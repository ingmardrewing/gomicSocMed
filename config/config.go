package config

import (
	"fmt"
	"os"
)

func GetPasswordHashForUser(user string) string {
	// TODO impl. user dependent hash retrieval
	return os.Getenv("GOMIC_REST_BASIC_AUTH_PASS")
}

func GetTwitterConsumerKey() string {
	s := os.Getenv("TWITTER_ORIGIN_CONSUMER_KEY")
	fmt.Println(s)
	return s
}

func GetTwitterConsumerSecret() string {
	s := os.Getenv("TWITTER_ORIGIN_CONSUMER_SECRET")
	fmt.Println(s)
	return s
}

func GetTwitterAccessToken() string {
	s := os.Getenv("TWITTER_ORIGIN_ACCESS_TOKEN")
	fmt.Println(s)
	return s
}

func GetTwitterAccessTokenSecret() string {
	s := os.Getenv("TWITTER_ORIGIN_ACCESS_TOKEN_SECRET")
	fmt.Println(s)
	return s
}

func GetTumblrConsumerKey() string {
	return os.Getenv("GOMIC_TUMBLR_CONSUMER_KEY")
}

func GetTumblrConsumerSecret() string {
	return os.Getenv("GOMIC_TUMBLR_CONSUMER_SECRET")
}

func GetTumblrToken() string {
	return os.Getenv("GOMIC_TUMBLR_TOKEN")
}

func GetTumblrTokenSecret() string {
	return os.Getenv("GOMIC_TUMBLR_TOKEN_SECRET")
}

func GetTags() []string {
	return []string{"#inked", "#inking", "#illustration", "#drawing", "#drawings", "#art", "#artwork", "#draw", "#painting", "#sketch", "#sketchbook", "#artist", "#comics", "#comicart", "#comic", "#graphicnovel", "#design", "#concept", "#conceptart", "#create", "#creative", "#image", "#imagination"}
}

func IsProd() bool {
	return os.Getenv("GOMIC_STAGE") == "prod"
}
