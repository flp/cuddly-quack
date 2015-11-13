package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

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

	http.Handle("/", &app.IndexHandler{})
	http.Handle("/drafts/{id}", &app.ShowHandler{})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
