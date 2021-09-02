package image_application

import (
	"io"

	image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"
)

type ImageUploader struct {
	UploaderProvider image_domain.ImageUploaderProvider
}

func (i *ImageUploader) Upload(image io.Reader) string {
	return i.UploaderProvider.UploadImage(image)
}
