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
	Link, ImgUrl, Title, TagsCsvString string
	Tags                               []string
}

func NewSocMedService() *restful.WebService {
	path := "/0.1/gomic/socmed"
	service := new(restful.WebService)
	service.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Rest base path: %s\n", path)

	service.Route(service.POST("/publish").Filter(basicAuthenticate).To(Publish))

	service.Route(service.POST("/tumblr/callback").To(TumblrCallback))

	service.Route(service.POST("/facebook/callback").To(FacebookCallback))
	service.Route(service.GET("/facebook/callback").To(FacebookCallback))

	service.Route(service.POST("/facebook/getAccessToken").To(FacebookInit))
	service.Route(service.GET("/facebook/getAccessToken").To(FacebookInit))

	return service
}

func basicAuthenticate(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	err := authenticate(request)
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
	return bcrypt.CompareHashAndPassword(stored_hash, given_pass)
}

func Publish(request *restful.Request, response *restful.Response) {
	c := new(Content)
	request.ReadEntity(c)
	err := checkContent(c)
	if err != nil {
		response.WriteErrorString(400, "400: Bad Request ("+err.Error()+")")
		return
	}

	p := prepareContent(c)

	//tweet(c)
	//postToTumblr(c)
	postToFacebook(c)

	response.WriteEntity(p)
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

	if len(msg) > 0 {
		return errors.New(strings.Join(msg, ", "))
	}
	return nil
}

func prepareContent(c *Content) *Content {
	if len(c.TagsCsvString) > 0 {
		c.Tags = strings.Split(c.TagsCsvString, ",")
	}
	return c
}
