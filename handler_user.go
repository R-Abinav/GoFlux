package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/R-Abinav/GoFlux/internal/database"
	"github.com/google/uuid"
	"github.com/R-Abinav/GoFlux/internal/auth"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body);

	params := parameters{};
	err := decoder.Decode(&params);

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err));
		return;
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	});

	if err != nil{
		respondWithError(w, 400,  fmt.Sprintf("Couldn't create user: %s", err));
		return 
	}

	respondWithJSON(w, 201, databaseUserToUser(user));
}

func (apiCfg *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request){
	apiKey, err := auth.GetApiKey(r.Header);
	if err != nil{
		respondWithError(w, 403, fmt.Sprintf("Auth Error: %s", err));
		return;
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey);
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %s", err));
	}

	respondWithJSON(w, 200, databaseUserToUser(user));
}