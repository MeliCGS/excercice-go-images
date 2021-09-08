package image_application

import image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"

type ImageCreator struct {
	Repo   image_domain.ImageRepository
	finder *image_domain.ImageFinder
}

func NewImageCreator(repo image_domain.ImageRepository) *ImageCreator {
	return &ImageCreator{
		Repo:   repo,
		finder: &image_domain.ImageFinder{Repo: repo},
	}
}

func (t *ImageCreator) Create(image *image_domain.Image) error {
	foundTodo, _ := t.finder.Find(image.ID)

	if foundTodo != nil {
		return image_domain.NewImageAlreadyExistsError(*image)
	}

	t.Repo.Add(image)
	return nil
}
