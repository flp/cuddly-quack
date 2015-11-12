package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ShowHandler struct{}

func (s *ShowHandler) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	room, ok := coordinators.DefaultInMemoryCoordinator[id]
	if !ok {
		rw.Write([]byte("Could not find room"))
		return
	}

	rw.Write([]byte(fmt.Sprintf("%v\n", room)))
}
