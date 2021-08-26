package image_application

import (
	image_domain "github.com/MeliCGS/excercice-go-images/api/savee-image/domain"
)

type ImageGetImage struct {
	Repo image_domain.ImageRepository
}

func (p *ImageGetImage) SearchImage() []*image_domain.Image {
	return p.Repo.GetListNumberRandom()
}
