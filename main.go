package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/flp/cuddly-quack/app"
	"github.com/flp/cuddly-quack/draft/coordinators"

func main() {
	coordinators.DefaultInMemoryCoordinator.CreateDraftRoom("test")

	// Seed the rng
	rand.Seed(time.Now().UnixNano())

	// Read BFZ data and generate a pack
	contents, _ := Asset("resources/BFZ.json")
	mtg.ReadCardData(contents)
	mpg := mtg.ModernPackGenerator{}
	p := mpg.GeneratePack()
	fmt.Println(p)

	http.Handle("/", &app.IndexHandler{})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
