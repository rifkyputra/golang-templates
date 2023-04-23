package modularHTTP

import (
	db "ModularHTTPGo/db"
	route "ModularHTTPGo/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
)

func StartServer(addr string) {

	m := mux.NewRouter()

	dbConn, err := db.CreateConnection()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer dbConn.Close()

	// Serve Swagger documentation
	swaggerUI := httpSwagger.WrapHandler
	m.PathPrefix("/swagger/").Handler(swaggerUI)

	// Serve Swagger JSON
	m.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swagger, err := swag.ReadDoc()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(swagger))
	})

	route.ArticleRouter(m, dbConn)
	route.AuthRouter(m, dbConn)
	route.UserRouter(m, dbConn)
	route.IndexRouter(m)

	server := http.Server{
		Addr:    addr,
		Handler: m,
	}
	log.Printf("Starting server on %s\n", addr)
	server.ListenAndServe()
}
