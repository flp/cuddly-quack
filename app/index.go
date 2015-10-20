package app

import (
	"net/http"
)

type IndexHandler struct{}

func (i *IndexHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hello world"))
}
