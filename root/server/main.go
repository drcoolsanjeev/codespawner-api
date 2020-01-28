package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-pg/pg/v9"

	"github.com/codespawner-api/root"
	"github.com/codespawner-api/root/postgres"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	} `json:"database"`
	Port string `json:"port"`
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

const defaultPort = "1401"

func main() {
	var config Config
	config, err := LoadConfiguration("config/conf.json")
	if err != nil {
		fmt.Println(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Port
	}
	// DB Config
	DB := postgres.New(&pg.Options{
		User:     config.Database.Username,
		Password: config.Database.Password,
		Database: config.Database.Dbname,
	})
	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})
	c := root.Config{Resolvers: &root.Resolver{
		UsersRepo: postgres.UserRepo{DB: DB},
	}}
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(root.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
