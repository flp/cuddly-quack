package main

import (
	"log"
	"net/http"

	"github.com/flp/cuddly-quack/app"
	"github.com/flp/cuddly-quack/draft/coordinators"
)

func main() {
	coordinators.DefaultInMemoryCoordinator.CreateDraftRoom("test")

	http.Handle("/", &app.IndexHandler{})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
