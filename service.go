package main

import (
	"errors"
	"log"
	"strings"

	fb "github.com/huandu/facebook"
	"golang.org/x/crypto/bcrypt"

	restful "github.com/emicklei/go-restful"
)

type Content struct {
	Link, ImgUrl, Title, TagsCsvString, Description string
	Tags                                            []string
}

func NewSocMedService() *restful.WebService {
	path := "/0.1/gomic/socmed"
	echo := "/echo"
	publish := "/all/publish"
	publishTwitter := "/twitter/publish"
	//publishFacebook := "/facebook/publish"
	publishTumblr := "/tumblr/publish"

	tumblrCallback := "/tumblr/callback"
	//facebookCallback := "/facebook/callback"
	//facebookGetAccessToken := "/facebook/getAccessToken"

	service := new(restful.WebService)
	service.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Adding POST route: %s\n", path+echo)
	service.Route(service.POST(echo).Filter(basicAuthenticate).To(Echo))

	log.Printf("Adding POST route: %s\n", path+publish)
	service.Route(service.POST(publish).Filter(basicAuthenticate).To(Publish))

	log.Printf("Adding POST route: %s\n", path+publishTwitter)
	service.Route(service.POST(publishTwitter).Filter(basicAuthenticate).To(PublishTwitter))

	/*
		log.Printf("Adding POST route: %s\n", path+publishFacebook)
		service.Route(service.POST(publishFacebook).Filter(basicAuthenticate).To(PublishFacebook))
	*/

	log.Printf("Adding POST route: %s\n", path+publishTumblr)
	service.Route(service.POST(publishTumblr).Filter(basicAuthenticate).To(PublishTumblr))

	log.Printf("Adding POST route: %s\n", path+tumblrCallback)
	service.Route(service.POST(tumblrCallback).To(TumblrCallback))

	/*
		log.Printf("Adding GET and POST route: %s\n", path+facebookCallback)
		service.Route(service.POST(facebookCallback).To(FacebookCallback))
		service.Route(service.GET(facebookCallback).To(FacebookCallback))

		log.Printf("Adding GET and POST route: %s\n", path+facebookGetAccessToken)
		service.Route(service.POST(facebookGetAccessToken).To(FacebookGetAccessToken))
		service.Route(service.GET(facebookGetAccessToken).To(FacebookGetAccessToken))
	*/

	return service
}

func basicAuthenticate(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	err := authenticate(request)
	log.Println(err)
	if err != nil {
		response.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		response.WriteErrorString(401, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(request, response)
}

func authenticate(req *restful.Request) error {
	user, pass, _ := req.Request.BasicAuth()
	log.Println("request from user: ", user)
	/*log.Println("pass: ", pass)*/
	given_pass := []byte(pass)

	existing_hash := env(GOMIC_BASIC_AUTH_PASS_HASH)
	stored_hash := []byte(existing_hash)

	/*
		cost := 4
		hash, err := bcrypt.GenerateFromPassword(given_pass, cost)
		if err != nil {
			log.Fatalln("error generating password:", err)
		}

		log.Println("existing hash:", existing_hash)
		log.Println("hashed new pw:", string(hash))
	*/

	return bcrypt.CompareHashAndPassword(stored_hash, given_pass)
}

func Echo(request *restful.Request, response *restful.Response) {
	err, c := readContent(request)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}
	response.WriteEntity(c)
}

func Publish(request *restful.Request, response *restful.Response) {
	err, c := readContent(request)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}
	tweet(c)
	postToTumblr(c)

	result := []fb.Result{}
	response.WriteEntity(result)
}

func PublishTwitter(request *restful.Request, response *restful.Response) {
	err, c := readContent(request)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}
	tweet_id := tweet(c)

	response.WriteEntity(tweet_id)
}

/*
func PublishFacebook(request *restful.Request, response *restful.Response) {
	err, c := readContent(request)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}
	result := postToFacebook(c)

	result := []fb.Result{}
	response.WriteEntity(result)
}
*/

func PublishTumblr(request *restful.Request, response *restful.Response) {
	err, c := readContent(request)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}
	postToTumblr(c)

	response.WriteEntity(c)
}

func checkContent(c *Content) error {
	msg := []string{}
	if len(c.Link) == 0 {
		msg = append(msg, "No Link given")
	}
	if len(c.ImgUrl) == 0 {
		msg = append(msg, "No ImgUrl given")
	}
	if len(c.Title) == 0 {
		msg = append(msg, "No Title given")
	}
	if len(c.TagsCsvString) == 0 {
		msg = append(msg, "No TagsCsvString given")
	}
	if len(c.Description) == 0 {
		msg = append(msg, "No Description given")
	}

	if len(msg) > 0 {
		return errors.New(strings.Join(msg, ", "))
	}
	return nil
}

func readContent(request *restful.Request) (error, *Content) {
	c := new(Content)
	request.ReadEntity(c)
	err := checkContent(c)
	if err != nil {
		return err, new(Content)
	}
	if len(c.TagsCsvString) > 0 {
		c.Tags = strings.Split(c.TagsCsvString, ",")
	}
	return nil, c
}
