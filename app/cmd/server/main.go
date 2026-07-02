package main

import (
	"devops-platform/internal/config"
	"devops-platform/internal/handlers"
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to DevOps Platform")
}

func main() {
	cfg := config.Load()
	http.HandleFunc("/", home)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/version", handlers.VersionHandler(cfg.AppVersion))

	log.Printf("Starting server on :%s", cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
