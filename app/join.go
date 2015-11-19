package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flp/cuddly-quack/draft/coordinators"
	"github.com/gorilla/mux"
)

type JoinHandler struct{}

func (s *JoinHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("Error registering user: %v", err)))
		return
	}

	var req JoinRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("Error registering user: %v", err)))
		return
	}

	err = coordinators.DefaultInMemoryCoordinator.RegisterUser(req.UserID, roomID)
	if err != nil {
		rw.Write([]byte(fmt.Sprintf("Error registering user: %v", err)))
		return
	}

	rw.Write([]byte("Successfully registered user for room"))
}

type JoinRequest struct {
	UserID string `json:"user_id"`
}
