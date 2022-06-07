package main

import (
	"fmt"
	myRouter "github.com/joemama0s/backend-developer-tests/rest-service/pkg/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("SP// Backend Developer Test - RESTful Service")
	fmt.Println()

	// TODO: Add RESTful web service here
	router := myRouter.NewPersonRouter()

	log.Fatal(http.ListenAndServe(":8000", router))

}
