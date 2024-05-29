package main

import (
	"fmt"
	"log"
	"net/http"
	"TWAMP/internal/handlers"
	"TWAMP/internal/commands"
	"TWAMP/internal/utils"
	"html/template"
)

func main() {
	// Create an instance of Handlers with the actual implementations
	h := &handlers.Handlers{
		StartServer:       commands.StartServer,
		ReadOutputFile:    utils.ReadOutputFile,
		GenerateChartData: utils.GenerateChartData,
		GenerateTableData: utils.GenerateTableData,
		ParseFiles:        template.ParseFiles,
	}

	fs := http.FileServer(http.Dir("app"))
	hn := http.StripPrefix("/app/", fs)
	http.Handle("/app/", hn)

	http.HandleFunc("/", h.Handler)
	http.HandleFunc("/submitted", h.Handler)
	http.HandleFunc("/result", h.Handler)

	port := 8080
	fmt.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
