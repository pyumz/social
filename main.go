package main

import (
	"fmt"
	"net/http"
	"time"
)

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

	fmt.Println("listening...")
	http.ListenAndServe(srv.Addr, nil)
	fmt.Println("Error is ", http.ListenAndServe(srv.Addr, nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{}\n"))
}
