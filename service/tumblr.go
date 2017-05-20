package service

import (
	"fmt"

	"github.com/MariaTerzieva/gotumblr"
	"github.com/ingmardrewing/gomicSocMed/config"
)

func postToTumblr(c *Content) {
	client := getTumblrClient()
	mappedContent := getMappedContent(c)

	if config.IsProd() {
		err := client.CreatePhoto(
			config.GetTumblrBlogName(),
			mappedContent)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getMappedContent(c *Content) map[string]string {
	return map[string]string{
		"link":    c.Link,
		"source":  c.ImgUrl,
		"caption": c.Title,
		"tags":    c.TagsCsvString,
		"state":   "published"}
}

func getTumblrClient() *gotumblr.TumblrRestClient {
	return gotumblr.NewTumblrRestClient(
		config.GetTumblrConsumerKey(),
		config.GetTumblrConsumerSecret(),
		config.GetTumblrToken(),
		config.GetTumblrTokenSecret(),
		config.GetTumblrCallbackUrl(),
		"http://api.tumblr.com")
}
