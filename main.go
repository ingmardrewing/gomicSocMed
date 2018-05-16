package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	restful "github.com/emicklei/go-restful"
	store "github.com/ingmardrewing/fsKeyValueStore"
	shared "github.com/ingmardrewing/gomicSocMedShared"
)

func main() {
	logfile := path.Join(shared.Env(shared.GOMIC_PATH_TO_LOG), "log")
	f, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	fmt.Println("Now logging to", logfile)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	store.Initialize(shared.Env(shared.PATH_TO_FILE_DB))
	restful.Add(NewSocMedService())
	port := shared.Env(shared.GOMIC_SOCMED_PROD_PORT)

	/*
		crt := shared.Env(shared.TLS_CERT_PATH)
		key := shared.Env(shared.TLS_KEY_PATH)
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
