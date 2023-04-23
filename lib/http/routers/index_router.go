package modularHTTP

import (
	service "ModularHTTPGo/services"
	. "ModularHTTPGo/types"

	"github.com/gorilla/mux"
)

func IndexRouter(m *mux.Router) Router {
	path := "/"

	get := m.HandleFunc(path, service.IndexHandler).Methods("GET")
	routes := []*mux.Route{get}

	return Router{Name: path, MuxRoute: routes}

}
