package app

import(
	"net/http"

	"github.com/gorilla/mux"
)

func (i *IndexHandler) ServeHttp(rw http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	room, ok := coordinators.DefaultInMemoryCoordinator[id]
	if !ok {
		rw.Write([]byte("Could not find room"))
		return
	}

	rw.Write([]byte(fmt.Sprintf("%v\n", room)))	
}
