package main

import (
	"log"
	"net/http"
	"os"

	image_infrastructure "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/infrastructure"
	"github.com/gorilla/mux"
)

// type Size struct {
// 	Width  int `json:"width"`
// 	Height int `json:"height"`
// }

// type ListRandom struct {
// 	Id     int `json:"id"`
// 	Height int `json:"height"`
// }

// type Image struct {
// 	ID    int    `json:"id"`
// 	Url   string `json:"url"`
// 	Autor string `json:"autor"`
// 	Alt   string `json:"alt"`
// 	Size  *Size  `json:"size"`
// }

// type IntRange struct {
// 	min, max int
// }

// func (ir *IntRange) NextRandom(r *rand.Rand) int {
// 	return r.Intn(ir.max-ir.min+1) + ir.min
// }

// func getHeight(total int) []*ListRandom {
// 	listHeight := []*ListRandom{}
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	intRange := IntRange{1, 10}
// 	intRangeId := IntRange{1, 1000}
// 	for i := 0; i < total; i++ {
// 		listHeight = append(listHeight, &ListRandom{
// 			Height: (intRange.NextRandom(r) * 100),
// 			Id:     intRangeId.NextRandom(r),
// 		})
// 	}
// 	return listHeight
// }

// func GetImageEndpoint(w http.ResponseWriter, r *http.Request) {
// 	total, err := strconv.Atoi(r.URL.Query().Get("total"))

// 	if err != nil {
// 		total = 10
// 	}

// 	list := getHeight(total)
// 	images := []*Image{}
// 	for i := 0; i < len(list); i++ {
// 		images = append(images, &Image{
// 			ID:    list[i].Id,
// 			Url:   fmt.Sprintf("https://picsum.photos/300/%v", list[i].Height),
// 			Size:  &Size{300, int(list[i].Height)},
// 			Autor: "prueba",
// 			Alt:   "prueb",
// 		})
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&images)
// }

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8100"
	}

	router := mux.NewRouter()
	router.HandleFunc("/images", image_infrastructure.GetImageHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
