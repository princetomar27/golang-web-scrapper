package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/princetomar27/golang_web_scrapper/rss/internal/auth"
	"github.com/princetomar27/golang_web_scrapper/rss/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	
	type parameters struct{
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON : %v",err))
		return
	}

	//  Now we have db access to create user
	user,err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil {
		respondWithError(w,401,fmt.Sprintf("Error creating user : %v",err))
		return
	}



	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
    apiKey, err := auth.GetAPIKeyFromHeaders(r.Header)

    if err != nil {
        respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
        return
    }
 
    user, err:= apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

	if err != nil {
        respondWithError(w, 404, fmt.Sprintf("User not found: %v", err))
        return
    }

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context());

	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Getting Error while fetching list of users : %v", err))
		return
	}

	respondWithJson(w, 200, users)
}

 