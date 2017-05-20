package service

import (
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

	return service
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	err := authenticate(req)
	if err != nil {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}

	chain.ProcessFilter(req, resp)
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
	p := prepareContent(c)

	response.WriteEntity(p)
}

func prepareContent(c *Content) *Content {
	if len(c.TagsCsvString) > 0 {
		c.Tags = strings.Split(c.TagsCsvString, ",")
	}
	return c
}
