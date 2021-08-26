package image_infrastructure

import (
	"math/rand"
	"time"

	image_domain "github.com/MeliCGS/excercice-go-images/api/savee-image/domain"
)

type IntRange struct {
	min, max int
}

func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func GetListNumberRandom(total int) []*image_domain.ListRandom {
	ListRandom := []*image_domain.ListRandom{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	intRange := IntRange{1, 10}
	intRangeId := IntRange{1, 1000}
	for i := 0; i < total; i++ {
		ListRandom = append(ListRandom, &image_domain.ListRandom{
			Height: (intRange.NextRandom(r) * 100),
			Id:     intRangeId.NextRandom(r),
		})
	}
	return ListRandom
}
