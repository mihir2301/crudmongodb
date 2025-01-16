package main

import (
	"fmt"
	"log"
	user "mihirproject/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome")
	r := user.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
