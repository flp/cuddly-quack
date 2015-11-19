package app

import (
	"fmt"
	"net/http"

	"github.com/flp/cuddly-quack/draft/coordinators"
	"github.com/gorilla/mux"
)

type ShowHandler struct{}

func (s *ShowHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	room, err := coordinators.DefaultInMemoryCoordinator.GetDraftRoom(id)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("Could not find room: %v", err)))
		return
	}

	rw.Write([]byte(fmt.Sprintf("%v\n", room)))
}
