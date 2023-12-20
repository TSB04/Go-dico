package main

import (
	"dico/dicoservice"
	"log"
	"net/http"
)

func main() {
	// Create a DicoService
	dicoService := dicoservice.NewDicoService("dico.txt")

	// Setup HTTP handlers
	http.HandleFunc("/add", dicoService.AddWordHandler)
	http.HandleFunc("/get", dicoService.GetWordHandler)
	http.HandleFunc("/list", dicoService.GetAllWordsHandler)
	http.HandleFunc("/remove", dicoService.RemoveWordHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8081", nil))

}
