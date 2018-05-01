package main

import "os"

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

func env(key string) string {
	return os.Getenv(key)
}

func IsProd() bool {
	return env(GOMIC_STAGE) == GOMIC_STAGE_PROD
}

func GetTlsPaths() (string, string) {
	cert := env(TLS_CERT_PATH)
	key := env(TLS_KEY_PATH)
	return cert, key
}
