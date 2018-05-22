package main

import (
	"fmt"
	"log"

	"github.com/MariaTerzieva/gotumblr"
	restful "github.com/emicklei/go-restful"
	shared "github.com/ingmardrewing/gomicSocMedShared"
)

func TumblrCallback(request *restful.Request, response *restful.Response) {
	// TODO
}

func postToTumblr(c *Content) {
	log.Println("Posting to tumblr")
	client := getTumblrClient()
	mappedContent := getMappedContent(c)

	if shared.IsProd() {
		err := client.CreatePhoto(
			shared.Env(shared.GOMIC_TUMBLR_BLOG_NAME),
			mappedContent)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Posting to tumblr succeeded")
		}
	}
}

func getMappedContent(c *Content) map[string]string {
	fmt.Println("Tumbler tags:", c.TagsCsvString)
	return map[string]string{
		"link":        c.Link,
		"source":      c.ImgUrl,
		"caption":     c.Title,
		"tags":        c.TagsCsvString,
		"description": c.Description,
		"state":       "published"}
}

func getTumblrClient() *gotumblr.TumblrRestClient {
	return gotumblr.NewTumblrRestClient(
		shared.Env(shared.GOMIC_TUMBLR_CONSUMER_KEY),
		shared.Env(shared.GOMIC_TUMBLR_CONSUMER_SECRET),
		shared.Env(shared.GOMIC_TUMBLR_TOKEN),
		shared.Env(shared.GOMIC_TUMBLR_TOKEN_SECRET),
		shared.Env(shared.GOMIC_TUMBLR_CALLBACK_URL),
		"http://api.tumblr.com")
}
