package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	json "encoding/json"

	"github.com/codespawner-api/root/api"
	"github.com/codespawner-api/root/models"
	"github.com/codespawner-api/root/server"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
)

type Config struct {
	Database struct {
		Host		string	`json:"host"`
		Port 		string	`json:"port"`
		Username	string	`json:"username"`
		Password	string	`json:"password"`
		Dbname		string	`json:"dbname"`
	} `json:"database"`
	Port	string	`json:"port"`	
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func main() {
	var config Config
	config, err := LoadConfiguration("config/conf.json")
	if err != nil {
		fmt.Println("%v", err)
	}
	router, db := initializeAPI(config)
	defer db.Close()

	fmt.Println("Serving Root on: 1401")
	
	log.Fatal(http.ListenAndServe(":" + config.Port, router))
}

func initializeAPI(config Config) (*chi.Mux, *models.Db) {
	// Create a new router
	router := chi.NewRouter()

	// Create a new connection to our pg database
	db, err := models.New(
		models.ConnString(
			config.Database.Host,
			config.Database.Port,
			config.Database.Username,
			config.Database.Password,
			config.Database.Dbname,
		),
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
