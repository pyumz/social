package main

import (
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	postUUID := strings.TrimPrefix(r.URL.Path, "/posts/")

	err := apiCfg.dbClient.DeletePost(postUUID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
