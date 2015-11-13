package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flp/cuddly-quack/draft/coordinators"
	"github.com/flp/cuddly-quack/wire"
)

type CreateDraftHandler struct{}

func (c *CreateDraftHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO Define JSON error bodies
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Printf("NO BODY")
		return
	}

	msg := &wire.CreateDraftRequest{}
	err = json.Unmarshal(data, msg)
	if err != nil {
		// TODO Define JSON error bodies
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Printf("err: %v", err)
		return
	}

	coordinators.DefaultInMemoryCoordinator.CreateDraftRoom(msg)
	rw.WriteHeader(http.StatusCreated)
}
