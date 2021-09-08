package image_domain

type ImageFinder struct {
	Repo ImageRepository
}

func (t *ImageFinder) Find(id int) (*Image, error) {
	image := t.Repo.Find(id)

	if image == nil {
		return nil, NewImageNotFoundError(id)
	}

	return image, nil
}
