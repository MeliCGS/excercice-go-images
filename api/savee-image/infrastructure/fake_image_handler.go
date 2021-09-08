package image_infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	image_application "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/application"
	image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"
	"github.com/gofiber/fiber/v2"
)

var mockMode = os.Getenv("MOCK") == "true"

type ImageHandler struct {
	creator  *image_application.ImageCreator
	searcher *image_application.ImageSearcher
	remover  *image_application.ImageRemover
	updater  *image_application.ImageUpdater
}

func NewImageHandler() *ImageHandler {
	var repo image_domain.ImageRepository

	repo = NewMongoImageRepository()

	creator := image_application.NewImageCreator(repo)
	searcher := image_application.NewImageSearcher(repo)
	remover := image_application.NewImageRemover(repo)
	updater := image_application.NewImageUpdater(repo)

	return &ImageHandler{
		creator:  creator,
		searcher: searcher,
		remover:  remover,
		updater:  updater,
	}
}

func (t *ImageHandler) GetAll(c *fiber.Ctx) error {
	return c.JSON(t.searcher.GetAll())
}

func (t *ImageHandler) Add(c *fiber.Ctx) error {
	image := &image_domain.Image{}

	c.BodyParser(image)
	err := t.creator.Create(image)

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

func (t *ImageHandler) Update(c *fiber.Ctx) error {
	image := &image_domain.Image{}

	c.BodyParser(image)
	err := t.updater.Update(image)

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

func (t *ImageHandler) Remove(c *fiber.Ctx) error {
	id := c.Params("id")
	iId, err := strconv.Atoi(id)

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	err = t.remover.Remove(iId)

	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(http.StatusOK).
		JSON(fiber.Map{"success": true})
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(2000)
	file, _, _ := r.FormFile("archivo")

	fakeImageController := image_application.ImageUploader{
		UploaderProvider: NewCloudinaryImageUploader(),
	}

	url := fakeImageController.Upload(file)

	fmt.Fprintf(w, url)

}

// func GetImageHandler(w http.ResponseWriter, r *http.Request) {
// 	total, err := strconv.Atoi(r.URL.Query().Get("total"))
// 	if err != nil {
// 		total = 10
// 	}

// 	fakeImageController := image_application.ImageGetImage{
// 		Repo: &FakeImageRepository{}, // revisar porque esta saliendo este error
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(fakeImageController.SearchImage(total))
// }
