package main

import (
	"net/http"
	"strings"
)

func (apiCfg apiConfig) handlerRetrievePosts(w http.ResponseWriter, r *http.Request) {
	userEmail := strings.TrimPrefix(r.URL.Path, "/posts/")

	posts, err := apiCfg.dbClient.GetPosts(userEmail)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
