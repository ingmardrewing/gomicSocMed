package service

import (
	"log"

	"github.com/MariaTerzieva/gotumblr"
	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicSocMed/config"
)

func TumblrCallback(request *restful.Request, response *restful.Response) {
	// TODO
}

func postToTumblr(c *Content) {
	log.Println("Posting to tumblr")
	client := getTumblrClient()
	mappedContent := getMappedContent(c)

	if config.IsProd() {
		err := client.CreatePhoto(
			config.GetTumblrBlogName(),
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
		config.GetTumblrConsumerKey(),
		config.GetTumblrConsumerSecret(),
		config.GetTumblrToken(),
		config.GetTumblrTokenSecret(),
		config.GetTumblrCallbackUrl(),
		"http://api.tumblr.com")
}
