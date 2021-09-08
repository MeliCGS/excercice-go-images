package image_domain

import "fmt"

type ImageNotFoundError struct {
	ID int
}

func NewImageNotFoundError(id int) error {
	return ImageNotFoundError{id}
}

func (t ImageNotFoundError) Error() string {
	return fmt.Sprintf("todo not found with id `%v`", t.ID)
}
