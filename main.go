package main

import (
	"net/http"
	"log"

	"github.com/flp/cuddly-quack/app"
)

func main() {
	http.Handle("/", &app.IndexHandler{})	
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
