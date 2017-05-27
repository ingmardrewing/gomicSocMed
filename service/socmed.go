package service

import (
	"errors"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicSocMed/config"
)

type Content struct {
	Link, ImgUrl, Title, TagsCsvString, Description string
	Tags                                            []string
}

func NewSocMedService() *restful.WebService {
	path := "/0.1/gomic/socmed"
	service := new(restful.WebService)
	service.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Rest base path: %s\n", path)

	service.Route(service.POST("/echo").Filter(basicAuthenticate).To(Echo))

	service.Route(service.POST("/publish").Filter(basicAuthenticate).To(Publish))
	service.Route(service.POST("/publishtwitter").Filter(basicAuthenticate).To(PublishTwitter))
	service.Route(service.POST("/publishfacebook").Filter(basicAuthenticate).To(PublishFacebook))
	service.Route(service.POST("/publishtumblr").Filter(basicAuthenticate).To(PublishTumblr))

	service.Route(service.POST("/tumblr/callback").To(TumblrCallback))

	service.Route(service.POST("/facebook/callback").To(FacebookCallback))
	service.Route(service.GET("/facebook/callback").To(FacebookCallback))

	service.Route(service.POST("/facebook/getAccessToken").To(FacebookInit))
	service.Route(service.GET("/facebook/getAccessToken").To(FacebookInit))

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
	given_pass := []byte(pass)
	stored_hash := []byte(config.GetPasswordHashForUser(user))
	//hash, _ := bcrypt.GenerateFromPassword(given_pass, coast)
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
	postToFacebook(c)

	response.WriteEntity(c)
}

func PublishTwitter(request *restful.Request, response *restful.Response) {
	err, c := readContent(request)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}
	tweet(c)

	response.WriteEntity(c)
}

func PublishFacebook(request *restful.Request, response *restful.Response) {
	err, c := readContent(request)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}
	postToFacebook(c)

	response.WriteEntity(c)
}

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
