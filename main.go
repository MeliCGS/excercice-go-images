package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Image struct {
	ID    int    `json:"id"`
	Url   string `json:"url"`
	Autor string `json:"autor"`
	Alt   string `json:"alt"`
	Size  *Size  `json:"size"`
}

type IntRange struct {
	min, max int
}

func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func getHeight(total int) []int {
	listHeight := []int{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intRange := IntRange{1, 10}
	for i := 0; i < total; i++ {
		listHeight = append(listHeight, intRange.NextRandom(r)*100)
	}
	return listHeight
}

func GetImageEndpoint(w http.ResponseWriter, r *http.Request) {
	total, err := strconv.Atoi(r.URL.Query().Get("total"))

	if err != nil {
		total = 10
	}

	list := getHeight(total)
	images := []*Image{}
	for i := 0; i < len(list); i++ {
		images = append(images, &Image{
			ID:    1,
			Url:   fmt.Sprintf("https://picsum.photos/300/%v", list[i]),
			Size:  &Size{300, int(list[i])},
			Autor: "prueba",
			Alt:   "prueb",
		})
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&images)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8100"
	}

	router := mux.NewRouter()
	router.HandleFunc("/images", GetImageEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
