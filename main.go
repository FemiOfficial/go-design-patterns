package main

import (
	servermiddleware "godesignpatterns/structural/decorator/server-middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", &servermiddleware.Server{
		LogWriter: os.Stdout,
		Handler: &servermiddleware.PlainServer{},
		Username: "someusername",
		Password: "somepassword",
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}