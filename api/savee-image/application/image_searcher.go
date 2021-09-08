package image_application

import image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"

type ImageSearcher struct {
	Repo image_domain.ImageRepository
}

func NewImageSearcher(repo image_domain.ImageRepository) *ImageSearcher {
	return &ImageSearcher{
		Repo: repo,
	}
}

func (t *ImageSearcher) GetAll() []*image_domain.Image {
	return t.Repo.GetAll()
}
