package main

import (
	"dico/dicoservice"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!"+"\n"+"Available endpoints:"+"\n"+"/add"+"\n"+"/get"+"\n"+"/list"+"\n"+"/remove")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(ds *dicoservice.DicoService) {

	// Call homePage at root endpoint
	http.HandleFunc("/", homePage)

	// Setup HTTP handlers
	http.HandleFunc("/add", ds.AddWordHandler)
	http.HandleFunc("/get", ds.GetWordHandler)
	http.HandleFunc("/list", ds.GetAllWordsHandler)
	http.HandleFunc("/remove", ds.RemoveWordHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {

	dicoService := dicoservice.NewDicoService("dico.txt")

	handleRequests(dicoService)
}
