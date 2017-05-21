package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicSocMed/db"
	"github.com/ingmardrewing/gomicSocMed/service"
)

func main() {
	db.Initialize()
	restful.Add(service.NewSocMedService())
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
