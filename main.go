package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"
	"database/sql"

	"github.com/R-Abinav/GoFlux/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
);

type apiConfig struct{
	DB *database.Queries
}

func main(){
	godotenv.Load()
	
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment or might be empty!")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not found in the environment or might be empty!")
	}

	//Connect to DB
	conn, err := sql.Open("postgres", dbUrl);
	if err != nil{
		log.Fatal("Cannot connect to database: ", err);
	}

	//Convert the type, we currently have sql.DB but we need type of database.Queries
	queries:= database.New(conn);

	apiCfg := apiConfig{
		DB: queries,
	}

	//New router object
	router := chi.NewRouter()

	//Cors Config
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins: []string{"https://*", "http://*"},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders: []string{"*"},
				ExposedHeaders: []string{"Link"},
				AllowCredentials: false, 
				MaxAge: 300,
			},
		),
	)

	v1Router := chi.NewRouter()
	
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}

