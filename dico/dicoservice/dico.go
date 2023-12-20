package dicoservice

import (
	"dico/dicorep"
	"encoding/json"
	"log"
	"net/http"
)

type DicoService struct {
	Dico *dicorep.Dico
}

// NewDicoService creates a new DicoService.
func NewDicoService(filename string) *DicoService {
	return &DicoService{
		Dico: dicorep.New(filename),
	}
}

// AddWordHandler handles the addition of a word and its definition.
func (ds *DicoService) AddWordHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request to add word")

    // Define a struct that matches the expected JSON structure
    type RequestData struct {
        Key   string `json:"key"`
        Value string `json:"value"`
    }

    // Decode the request body into the struct
    decoder := json.NewDecoder(r.Body)
    var requestData RequestData
    err := decoder.Decode(&requestData)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Use the data from the request in your logic
    ds.Dico.Add(requestData.Key, requestData.Value)
    w.WriteHeader(http.StatusCreated)
}

// GetWordHandler handles getting the definition of a word.
func (ds *DicoService) GetWordHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get word")
	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing 'word' parameter", http.StatusBadRequest)
		return
	}

	definition := ds.Dico.Couple[word]
	if definition != "" {
		response := map[string]string{word: definition}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Word not found", http.StatusNotFound)
	}
}

// GetWordHandler handles getting all the words and their definitions.
func (ds *DicoService) GetAllWordsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get all words")
	json.NewEncoder(w).Encode(ds.Dico.Couple)
}


// RemoveWordHandler handles the removal of a word and its definition.
func (ds *DicoService) RemoveWordHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to remove word")
	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Missing 'word' parameter", http.StatusBadRequest)
		return
	}

	ds.Dico.Remove(word)
	w.WriteHeader(http.StatusNoContent)
}
