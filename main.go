package main

import (
	"fmt"
	"log"
	"net/http"
	"TWAMP/internal/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("app"))
	hn := http.StripPrefix("/app/", fs)
	http.Handle("/app/", hn)

	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/submited", handlers.FormHandler)

	port := 8080
	fmt.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

