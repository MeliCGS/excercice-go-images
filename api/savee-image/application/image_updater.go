package image_application

import image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"

type ImageUpdater struct {
	Repo   image_domain.ImageRepository
	finder *image_domain.ImageFinder
}

func NewImageUpdater(repo image_domain.ImageRepository) *ImageUpdater {
	return &ImageUpdater{
		Repo:   repo,
		finder: &image_domain.ImageFinder{Repo: repo},
	}
}

func (t *ImageUpdater) Update(image *image_domain.Image) error {
	// Verify image existance
	_, err := t.finder.Find(image.ID)

	if err != nil {
		return err
	}

	t.Repo.Update(image)
	return nil
}
