package image_infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	image_application "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/application"
)

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(2000)
	file, _, _ := r.FormFile("archivo")
	// f, err := os.OpenFile("./files/"+fileInfo.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	/*if err != nil {
		log.Fatal(err)
		return
	}

	defer f.Close()*/

	fakeImageController := image_application.ImageUploader{
		UploaderProvider: NewCloudinaryImageUploader(),
	}

	url := fakeImageController.Upload(file)

	// io.Copy(f, file)

	fmt.Fprintf(w, url)

}

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
