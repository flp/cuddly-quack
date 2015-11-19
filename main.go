package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/flp/cuddly-quack/app"
	"github.com/flp/cuddly-quack/draft/coordinators"
	"github.com/flp/cuddly-quack/mtg"
	"github.com/flp/cuddly-quack/wire"
)

func main() {
	coordinators.DefaultInMemoryCoordinator.CreateDraftRoom(&wire.CreateDraftRequest{
		Name: "test",
		Set:  "BFZ",
	})

	// Seed the rng
	rand.Seed(time.Now().UnixNano())

	// Generate a pack
	mpg := mtg.ModernPackGenerator{}
	p, _ := mpg.GeneratePack(mtg.BattleForZendikar)
	fmt.Println(p)

	r := mux.NewRouter()

	r.Handle("/", &app.IndexHandler{}).Methods("GET")
	r.Handle("/drafts/{id}", &app.ShowHandler{}).Methods("GET")
	r.Handle("/drafts", &app.CreateDraftHandler{}).Methods("POST")
	r.Handle("/drafts/{room_id}/join", &app.JoinHandler{}).Methods("POST")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
