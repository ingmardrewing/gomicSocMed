package main

import (
	"log"

	"github.com/MariaTerzieva/gotumblr"
	restful "github.com/emicklei/go-restful"
)

func TumblrCallback(request *restful.Request, response *restful.Response) {
	// TODO
}

func postToTumblr(c *Content) {
	log.Println("Posting to tumblr")
	client := getTumblrClient()
	mappedContent := getMappedContent(c)

	if IsProd() {
		err := client.CreatePhoto(
			env(GOMIC_TUMBLR_BLOG_NAME),
			mappedContent)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Posting to tumblr succeeded")
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
		env(GOMIC_TUMBLR_CONSUMER_KEY),
		env(GOMIC_TUMBLR_CONSUMER_SECRET),
		env(GOMIC_TUMBLR_TOKEN),
		env(GOMIC_TUMBLR_TOKEN_SECRET),
		env(GOMIC_TUMBLR_CALLBACK_URL),
		"http://api.tumblr.com")
}
