package transfer

import "github.com/gorilla/mux"

func Router() *mux.Router {
	r := mux.NewRouter()
	return r
}
