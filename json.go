package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string){
	// if status code is >= 500
	if(code >= 500){
		log.Printf("Got 5XX error: %v", message)
	} 

	type errResponse struct{
		Error string `json:"error"`
	}

	respondWithJson(w,code, errResponse{
		Error: message,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){

	// 1. Marshal payload into JSON 
	jsonResponse, err := json.Marshal(payload)

	if err != nil{
		fmt.Printf("Marshal error: %v\n", err)
		w.WriteHeader(500)
		return
	}

	// 2. Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Print(string(jsonResponse))
	w.Write(jsonResponse)
}
