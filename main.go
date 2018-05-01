package main

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	store "github.com/ingmardrewing/fsKeyValueStore"
)

func main() {
	restful.Add(NewSocMedService())
	store.Initialize()

	port := "8443"

	/*
		port := "443"
		log.Println("Reading crt and key data from files:")
		crt, key := config.GetTlsPaths()

		log.Println("Path to crt file: " + crt)
		log.Println("Path to key file: " + key)
		log.Println("Starting to serve via TLS on Port: " + port)

		err := http.ListenAndServeTLS(":"+port, crt, key, nil)
	*/
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
