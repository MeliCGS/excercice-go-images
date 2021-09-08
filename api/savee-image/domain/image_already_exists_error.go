package image_domain

import "fmt"

type ImageAlreadyExistsError struct {
	ID Image
}

func NewImageAlreadyExistsError(id Image) error {
	return ImageAlreadyExistsError{id}
}

func (t ImageAlreadyExistsError) Error() string {
	return fmt.Sprintf("todo with id `%f` already exists", t.ID)
}
