package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"stockbit-backend/http/handler"
	"stockbit-backend/http/middlewares"

	"github.com/gocraft/dbr"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type App struct {
	Router *mux.Router
}

func RunServer(dbConn *dbr.Connection) {
	port := viper.GetString("PORT")
	r := mux.NewRouter()
	r.Use(middlewares.LoggingMiddleware)

	getMovies := handler.GetMovies{DBConn: dbConn}
	getMovie := handler.GetMovie{DBConn: dbConn}
	r.Handle("/movies", getMovies).
		Queries(
			"q", "{q}",
			"page", "{page}").Methods(http.MethodGet)
	r.Handle("/movie/{id}", getMovie).Methods(http.MethodGet)

	fmt.Printf("\n Server starting on Port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CombinedLoggingHandler(os.Stdout, r)))
}
