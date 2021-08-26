package image_domain

type ImageRepository interface {
	GetImageEndpoint(total int) []*Image
}
