package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Image struct {
	ID    int    `json:"id"`
	Url   string `json:"url"`
	Autor string `json:"autor"`

	size struct {
		width  int
		height int
	}

	Alt string `json: "Alt"`
}

func Images() {

}

func GetImageEndpoint(w http.ResponseWriter, r http.Request) {
	json.NewEncoder(w).Encode(images)
}

func main() {
	router := mux.NewRouter()
}
