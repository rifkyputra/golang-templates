package modularHTTP

import (
	. "ModularHTTPGo/middlewares"
	. "ModularHTTPGo/services"
	. "ModularHTTPGo/types"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func ArticleRouter(m *mux.Router, db *sql.DB) Router {
	path := "/article"
	get := m.HandleFunc(path, GetHandlerBranch(db)).Methods("GET")
	post := m.Handle(path, AuthMiddleware(http.HandlerFunc(PostArticleHandler(db)))).Methods("POST")
	put := m.HandleFunc(path, UpdateArticleHandler(db)).Methods("PUT")
	delete := m.Handle(path, AuthMiddleware(http.HandlerFunc(DeleteArticleHandler(db)))).Methods("DELETE")

	routes := []*mux.Route{get, post, put, delete}

	return Router{
		Name:     path,
		MuxRoute: routes,
	}
}

func GetHandlerBranch(db *sql.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Query().Has("id") {
			GetArticleByIdHandler(db)(w, r)
			return
		}

		GetArticleHandler(db)(w, r)
	}

}
