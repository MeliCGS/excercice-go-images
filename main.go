package main

import (
	"os"

	image_infrastructure "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/infrastructure"
	//"github.com/gorilla/mux"
	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8100"
	}

	newImageHandler := image_infrastructure.NewImageHandler()

	//router := mux.NewRouter()
	//router.HandleFunc("/images", image_infrastructure.GetImageHandler).Methods("GET")
	//router.HandleFunc("/images", image_infrastructure.UploadImageHandler).Methods("PUT")
	//router.HandleFunc("/images", newImageHandler.GetAll).Methods("GET")

	router := fiber.New()

	router.Post("/images", newImageHandler.Add)
	router.Get("/images", newImageHandler.GetAll)
	router.Patch("/images", newImageHandler.Update)
	router.Delete("/images/:id", newImageHandler.Remove)

	router.Listen(":" + port)
}
