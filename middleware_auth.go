package main

import (
	"fmt"
	"net/http"

	"github.com/princetomar27/golang_web_scrapper/rss/internal/auth"
	"github.com/princetomar27/golang_web_scrapper/rss/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User) 

func (cfg *apiConfig) middlewareAuth (handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetAPIKeyFromHeaders(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
	 
		user, err:= cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	
	}
}