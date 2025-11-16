package main

import (
	"net/http"
	"fmt"
	"github.com/R-Abinav/GoFlux/internal/database"
	"github.com/R-Abinav/GoFlux/internal/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header);
		if err != nil{
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %s", err));
			return;
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey);
		if err != nil{
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %s", err));
		}

		handler(w, r, user);
	}
}