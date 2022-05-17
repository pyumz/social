package main

import (
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")

	user, err := apiCfg.dbClient.GetUser(userEmail)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
