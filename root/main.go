package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codespawner-api/root/api"
	"github.com/codespawner-api/root/models"
	"github.com/codespawner-api/root/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

func main() {
	router, db := initializeAPI()
	defer db.Close()
	fmt.Println("Serving Root on: 1401")
	
	log.Fatal(http.ListenAndServe(":1401", router))
}

func initializeAPI() (*chi.Mux, *models.Db) {
	// Create a new router
	router := chi.NewRouter()

	// Create a new connection to our pg database
	db, err := models.New(
		models.ConnString("localhost", 5432, "iamsaquib", "90627", "codespawner"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create our root query for graphql
	rootQuery := api.NewRoot(db)
	// Create a new graphql schema, passing in the the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well
	// as the address of our graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}

	// Add some middleware to our router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,          // log api request calls
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,       // recover from panics without crashing server
	)

	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())

	return router, db
}
