package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir("./app")))
	mux.Handle("/app/", appHandler)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/app/\n", port)
	log.Fatal(srv.ListenAndServe())

}
