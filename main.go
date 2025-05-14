package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/paultustain/projectIdea/m/v2/internal/database"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	const port = "8080"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		db: dbQueries,
	}
	fmt.Println(apiCfg)
	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir("./app")))
	mux.Handle("/app/", appHandler)

	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/app/\n", port)
	log.Fatal(srv.ListenAndServe())

}
