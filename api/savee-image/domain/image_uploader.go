package image_domain

import "io"

type ImageUploaderProvider interface {
	UploadImage(image io.Reader) (url string)
}
