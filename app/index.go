package app

import (
	"fmt"
	"net/http"

	"github.com/flp/cuddly-quack/draft/coordinators"
)

type IndexHandler struct{}

func (i *IndexHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// TODO settle on some kind of templating engine.
	for _, room := range coordinators.DefaultInMemoryCoordinator.DraftRooms {
		rw.Write([]byte(fmt.Sprintf("%v\n", room)))
	}
}
