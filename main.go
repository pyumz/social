package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	database "github.com/pyumz/social/internal"
)

type errorBody struct {
	Error string `json:"error"`
}

type apiConfig struct {
	dbClient database.Client
}

func main() {
	m := http.NewServeMux()

	dbClient := database.NewClient("db.json")
	err := dbClient.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		dbClient: dbClient,
	}

	m.HandleFunc("/users", apiCfg.endpointUsersHandler)
	m.HandleFunc("/users/", apiCfg.endpointUsersHandler)

	m.HandleFunc("/posts", apiCfg.endpointPostsHanlder)
	m.HandleFunc("/posts/", apiCfg.endpointPostsHanlder)

	const addr = "localhost:8081"

	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("listening on", srv.Addr)
	srv.ListenAndServe()
	fmt.Println("Error is ", srv.ListenAndServe())
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, errorBody{
		Error: err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		response, _ := json.Marshal(errorBody{
			Error: "Error marshalling womp",
		})
		w.Write(response)
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}
