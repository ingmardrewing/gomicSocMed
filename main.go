package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicSocMed/config"
	"github.com/ingmardrewing/gomicSocMed/db"
	"github.com/ingmardrewing/gomicSocMed/service"
)

func main() {
	db.Initialize()
	restful.Add(service.NewSocMedService())

	crt, key := config.GetTlsPaths()
	log.Printf("crt path: %s\n", crt)
	log.Printf("key path: %s\n", key)
	err := http.ListenAndServeTLS(":443", crt, key, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
