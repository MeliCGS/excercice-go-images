package image_application

import (
	image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"
)

type ImageGetImage struct {
	Repo image_domain.ImageRepository
}

func (i *ImageGetImage) SearchImage(total int) []*image_domain.Image {
	return i.Repo.GetImageEndpoint(total)
}
