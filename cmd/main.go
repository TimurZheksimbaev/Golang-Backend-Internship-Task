package main

import (
	"go-backend-internship/routers"
	"log"
	"net/http"
)


func main() {
	router := routers.GetRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}