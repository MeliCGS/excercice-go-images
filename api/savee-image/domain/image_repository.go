package image_domain

type ImageRepository interface {
	Add(image *Image)
	GetAll() []*Image
	Remove(id int)
	Update(image *Image)
	Find(id int) *Image
	//GetImageEndpoint(total int) []*Image
}
