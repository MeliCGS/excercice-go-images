package image_infrastructure

import (
	"fmt"
	"math/rand"
	"time"

	image_domain "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/domain"
)

type IntRange struct {
	min, max int
}
type ListRandom struct {
	Id     int `json:"id"`
	Height int `json:"height"`
}

type FakeImageRepository struct{}

func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func GetListNumberRandom(total int) []*ListRandom {
	List := []*ListRandom{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intRange := IntRange{1, 10}
	intRangeId := IntRange{1, 1000}
	for i := 0; i < total; i++ {
		List = append(List, &ListRandom{
			Height: (intRange.NextRandom(r) * 100),
			Id:     intRangeId.NextRandom(r),
		})
	}
	return List
}

func (f *FakeImageRepository) GetImageEndpoint(total int) []*image_domain.Image {
	list := GetListNumberRandom(total)
	images := []*image_domain.Image{}
	for i := 0; i < len(list); i++ {
		images = append(images, &image_domain.Image{
			ID:    list[i].Id,
			Url:   fmt.Sprintf("https://picsum.photos/300/%v", list[i].Height),
			Size:  &image_domain.Size{300, int(list[i].Height)},
			Autor: "prueba",
			Alt:   "prueb",
		})
	}
	return images
}
