package main

import (
	"log"
	"net/http"

	"github.com/flp/cuddly-quack/app"
)

func main() {
	http.Handle("/", &app.IndexHandler{})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
