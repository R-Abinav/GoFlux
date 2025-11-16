package main

import "net/http"

type HealthResponse struct{
	Ok bool `json:"ok"`
	Message string `json:"message"`
}

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, 200, HealthResponse{
		Ok: true,
		Message: "GoFlux API is up and running!",
	});
}