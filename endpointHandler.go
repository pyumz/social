package main

import (
	"errors"
	"net/http"
)

func (apiCfg apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		apiCfg.handlerCreateUser(w, r)
	default:
		respondWithError(w, 404, errors.New("Unsupported API method call"))
	}
}
