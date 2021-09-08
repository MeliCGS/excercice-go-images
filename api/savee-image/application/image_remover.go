package image_application

import (
	image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"
)

type ImageRemover struct {
	Repo   image_domain.ImageRepository
	finder *image_domain.ImageFinder
}

func NewImageRemover(repo image_domain.ImageRepository) *ImageRemover {
	return &ImageRemover{
		Repo:   repo,
		finder: &image_domain.ImageFinder{Repo: repo},
	}
}

func (t *ImageRemover) Remove(id int) error {
	// Verify todo existance
	_, err := t.finder.Find(id)

	if err != nil {
		return err
	}

	t.Repo.Remove(id)
	return nil
}
