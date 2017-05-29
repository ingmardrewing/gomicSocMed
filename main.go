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

	port := "443"

	log.Println("Reading crt and key data from files:")
	crt, key := config.GetTlsPaths()

	log.Println("Path to crt file: " + crt)
	log.Println("Path to key file: " + key)
	log.Println("Starting to serve via TLS on Port: " + port)

	err := http.ListenAndServeTLS(":"+port, crt, key, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
