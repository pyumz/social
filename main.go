package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	database "github.com/pyumz/social/internal"
)

type errorBody struct {
	Error string `json:"error"`
}

func main() {

	m := http.NewServeMux()
	const addr = "localhost:8081"

	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("HandleFunc method called ")
	http.HandleFunc("/", testHandler)

	fmt.Println("HandleFunc Error method called")
	http.HandleFunc("/err", testErrHandler)

	fmt.Println("listening...")
	http.ListenAndServe(srv.Addr, nil)
	fmt.Println("Error is ", http.ListenAndServe(srv.Addr, nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, database.User{
		Email: "test@browser.com",
	})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("Server error"))
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
