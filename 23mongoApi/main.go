package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aman0106/ApiWithMongo/router"
)

func main() {
	fmt.Println("Main Program started")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":8000", r))
}
