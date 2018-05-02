package main

import (
	"log"
	"net/http"
	"os"
	"path"

	restful "github.com/emicklei/go-restful"
	store "github.com/ingmardrewing/fsKeyValueStore"
)

func main() {
	logfile := path.Join(env(GOMIC_PATH_TO_LOG), "log")
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

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
	log.Println("Listening on port", port, "via regular http")
	err = http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}
