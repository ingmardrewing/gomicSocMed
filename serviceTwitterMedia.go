package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	shared "github.com/ingmardrewing/gomicSocMedShared"
)

type PhotoTweet struct {
	MediaIds uint64 `json:"media_ids"`
	Status   string `json:"status"`
}

func getAnacondaApi() *anaconda.TwitterApi {
	return anaconda.NewTwitterApiWithCredentials(
		shared.Env(shared.TWITTER_REPEAT_ACCESS_TOKEN),
		shared.Env(shared.TWITTER_REPEAT_ACCESS_TOKEN_SECRET),
		shared.Env(shared.TWITTER_REPEAT_CONSUMER_KEY),
		shared.Env(shared.TWITTER_REPEAT_CONSUMER_SECRET))
}

func tweetMedia(c *shared.Content) anaconda.Tweet {
	api := getAnacondaApi()
	mediaId := uploadMedia(api, c.ImgUrl)

	v := url.Values{}
	v.Set("media_ids", strconv.FormatInt(mediaId, 10))

	result, err := api.PostTweet(getMediaTweetText(c), v)
	if err != nil {
		log.Println(err)
	}
	return result
}

func uploadMedia(api *anaconda.TwitterApi, url string) int64 {
	image := getPngFromUrl(url)
	imageBase64 := base64.StdEncoding.EncodeToString(image)

	mediaResponse, err := api.UploadMedia(imageBase64)
	if err != nil {
		fmt.Println(err)
	}
	return mediaResponse.MediaID
}

func getPngFromUrl(url string) []byte {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Trouble reading JSON response body!")
	}
	return contents
}

func getMediaTweetText(c *shared.Content) string {
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
