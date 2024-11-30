package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/princetomar27/golang_web_scrapper/rss/internal/database"
	"github.com/rs/cors"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main(){ go
	fmt.Println("Hey Prince")
	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == ""{
		log.Fatal("Empty port string!")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == ""{
		log.Fatal("Empty DB_URL!")
	}

	conn, err := sql.Open("postgres",dbURL)
	if err != nil{
		log.Fatal("Error connecting to database", err)
	}
 
	apiCfg := apiConfig{
		DB: database.New(conn) ,
	}
	 

	fmt.Println("Port: ", portString)

	router := chi.NewRouter()

	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://*", "https://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "*"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
    })



	router.Use(c.Handler)

	v1Router := chi.NewRouter()
	v1Router.HandleFunc("/ready", handlerReadiness)
	v1Router.Get("/err", handleError)
	v1Router.Post("/users",apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.handleGetUserByAPIKey)
	v1Router.Get("/users", apiCfg.handleGetAllUsers)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}


	fmt.Println("Server listening on PORT :", portString)
	if err := server.ListenAndServe();
	err != nil{
		log.Fatal(err)
	}

}