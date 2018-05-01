package config

import (
	"fmt"
	"os"
)

const (
	GOMIC_BASIC_AUTH_PASS_HASH         = "GOMIC_BASIC_AUTH_PASS_HASH"
	TWITTER_REPEAT_CONSUMER_KEY        = "TWITTER_REPEAT_CONSUMER_KEY"
	TWITTER_REPEAT_CONSUMER_SECRET     = "TWITTER_REPEAT_CONSUMER_SECRET"
	TWITTER_REPEAT_ACCESS_TOKEN        = "TWITTER_REPEAT_ACCESS_TOKEN"
	TWITTER_REPEAT_ACCESS_TOKEN_SECRET = "TWITTER_REPEAT_ACCESS_TOKEN_SECRET"
	TWITTER_ORIGIN_CONSUMER_KEY        = "TWITTER_ORIGIN_CONSUMER_KEY"
	TWITTER_ORIGIN_CONSUMER_SECRET     = "TWITTER_ORIGIN_CONSUMER_SECRET"
	TWITTER_ORIGIN_ACCESS_TOKEN        = "TWITTER_ORIGIN_ACCESS_TOKEN"
	TWITTER_ORIGIN_ACCESS_TOKEN_SECRET = "TWITTER_ORIGIN_ACCESS_TOKEN_SECRET"
	GOMIC_TUMBLR_CALLBACK_URL          = "GOMIC_TUMBLR_CALLBACK_URL"
	GOMIC_TUMBLR_BLOG_NAME             = "GOMIC_TUMBLR_BLOG_NAME"
	GOMIC_TUMBLR_CONSUMER_KEY          = "GOMIC_TUMBLR_CONSUMER_KEY"
	GOMIC_TUMBLR_CONSUMER_SECRET       = "GOMIC_TUMBLR_CONSUMER_SECRET"
	GOMIC_TUMBLR_TOKEN                 = "GOMIC_TUMBLR_TOKEN"
	GOMIC_TUMBLR_TOKEN_SECRET          = "GOMIC_TUMBLR_TOKEN_SECRET"
	FB_DEVABODE_ID                     = "FB_DEVABODE_ID"
	FB_APPLICATION_ID                  = "FB_APPLICATION_ID"
	FB_APPLICATION_SECRET              = "FB_APPLICATION_SECRET"
	FB_CALLBACK_URL                    = "FB_CALLBACK_URL"
	GOMIC_STAGE                        = "GOMIC_STAGE"
	GOMIC_STAGE_PROD                   = "prod"
	DB_GOMIC_USER                      = "DB_GOMIC_USER"
	DB_GOMIC_PASS                      = "DB_GOMIC_PASS"
	DB_GOMIC_NAME                      = "DB_GOMIC_NAME"
	DB_GOMIC_HOST                      = "DB_GOMIC_HOST"
	TLS_CERT_PATH                      = "TLS_CERT_PATH"
	TLS_KEY_PATH                       = "TLS_KEY_PATH"
)

func GetPasswordHashForUser(user string) string {
	// TODO impl. user dependent hash retrieval
	return os.Getenv(GOMIC_BASIC_AUTH_PASS_HASH)
}

func GetTwitterRepeatConsumerKey() string {
	return os.Getenv(TWITTER_REPEAT_CONSUMER_KEY)
}

func GetTwitterRepeatConsumerSecret() string {
	return os.Getenv(TWITTER_REPEAT_CONSUMER_SECRET)
}

func GetTwitterRepeatAccessToken() string {
	return os.Getenv(TWITTER_REPEAT_ACCESS_TOKEN)
}

func GetTwitterRepeatAccessTokenSecret() string {
	return os.Getenv(TWITTER_REPEAT_ACCESS_TOKEN_SECRET)
}

func GetTwitterConsumerKey() string {
	return os.Getenv(TWITTER_ORIGIN_CONSUMER_KEY)
}

func GetTwitterConsumerSecret() string {
	return os.Getenv(TWITTER_ORIGIN_CONSUMER_SECRET)
}

func GetTwitterAccessToken() string {
	return os.Getenv(TWITTER_ORIGIN_ACCESS_TOKEN)
}

func GetTwitterAccessTokenSecret() string {
	return os.Getenv(TWITTER_ORIGIN_ACCESS_TOKEN_SECRET)
}

func GetTumblrCallbackUrl() string {
	return os.Getenv(GOMIC_TUMBLR_CALLBACK_URL)
}

func GetTumblrBlogName() string {
	return os.Getenv(GOMIC_TUMBLR_BLOG_NAME)
}

func GetTumblrConsumerKey() string {
	return os.Getenv(GOMIC_TUMBLR_CONSUMER_KEY)
}

func GetTumblrConsumerSecret() string {
	return os.Getenv(GOMIC_TUMBLR_CONSUMER_SECRET)
}

func GetTumblrToken() string {
	return os.Getenv(GOMIC_TUMBLR_TOKEN)
}

func GetTumblrTokenSecret() string {
	return os.Getenv(GOMIC_TUMBLR_TOKEN_SECRET)
}

func GetFacebookPageId() string {
	return os.Getenv(FB_DEVABODE_ID)
}

func GetFacebookApplicationId() string {
	return os.Getenv(FB_APPLICATION_ID)
}

func GetFacebookApplicationSecret() string {
	return os.Getenv(FB_APPLICATION_SECRET)
}

func GetFacebookCallbackUrl() string {
	return os.Getenv(FB_CALLBACK_URL)
}

func IsProd() bool {
	return os.Getenv(GOMIC_STAGE) == GOMIC_STAGE_PROD
}

func GetDsn() string {
	user := os.Getenv(DB_GOMIC_USER)
	pass := os.Getenv(DB_GOMIC_PASS)
	name := os.Getenv(DB_GOMIC_NAME)
	host := os.Getenv(DB_GOMIC_HOST)
	return fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)
}

func GetTlsPaths() (string, string) {
	cert := os.Getenv(TLS_CERT_PATH)
	key := os.Getenv(TLS_KEY_PATH)
	return cert, key
}
