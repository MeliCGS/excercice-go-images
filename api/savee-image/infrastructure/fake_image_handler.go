package image_infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"

	image_application "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/application"
)

func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	total, err := strconv.Atoi(r.URL.Query().Get("total"))
	if err != nil {
		total = 10
	}
	fakeImageController := image_application.ImageGetImage{
		Repo: &FakeImageRepository{},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fakeImageController.SearchImage(total))
}
