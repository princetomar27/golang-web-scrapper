package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


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
	w.Write(jsonResponse)
}