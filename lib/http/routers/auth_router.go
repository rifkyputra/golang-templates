package modularHTTP

import (
	. "ModularHTTPGo/services"
	. "ModularHTTPGo/types"
	"database/sql"

	"github.com/gorilla/mux"
)

func AuthRouter(m *mux.Router, db *sql.DB) Router {
	path := "/auth"

	post := m.HandleFunc(path, PostAuthHandler(db))

	return Router{
		Name: path,
		MuxRoute: []*mux.Route{
			post,
		},
	}
}
