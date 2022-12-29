package main

import (
	"log"
	"net/http"
	"os"

	modules "github.com/jaimeag/skims-takehome/modules"
)

func main() {
	logger := log.New(os.Stdout, "skims-takehome", log.LstdFlags)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: modules.NewRouter(),
	}
	logger.Print("Starting server on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalf("Failed to start server %v\n", err)
	}
}
