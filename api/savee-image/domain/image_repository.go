package image_domain

type ImageRepository interface {
	GetListNumberRandom() []*ListRandom
}
