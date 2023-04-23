package modularHTTP

import (
	"database/sql"

	. "ModularHTTPGo/services"
	. "ModularHTTPGo/types"

	"github.com/gorilla/mux"
)

func UserRouter(m *mux.Router, db *sql.DB) Router {
	path := "/user"

	get := m.HandleFunc(path, GetUserHandler(db)).Methods("GET")
	post := m.HandleFunc(path, PostUserHandler(db)).Methods("POST")
	put := m.HandleFunc(path, UpdateUserHandler(db)).Methods("PUT")
	delete := m.HandleFunc(path, DeleteArticleHandler(db)).Methods("DELETE")

	routes := []*mux.Route{get, post, put, delete}

	return Router{
		Name:     path,
		MuxRoute: routes,
	}
}
