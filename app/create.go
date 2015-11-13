package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flp/cuddly-quack/draft/coordinators"
	"github.com/flp/cuddly-quack/wire"
)

type CreateDraftHandler struct{}

func (c *CreateDraftHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// 400 bad request
	}

	var msg *wire.CreateDraftRequest
	err = json.Unmarshal(data, msg)
	if err != nil {
		// 400 bad request
	}

	coordinators.DefaultInMemoryCoordinator.CreateDraftRoom(msg)
	// 201 Created
}
