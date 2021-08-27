package image_infrastructure_test

import (
	"testing"

	fake_image_repository "github.com/MeliCGS/NewRepoGoSavee/api/savee-image/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestGetList(t *testing.T) {
	fakeImage := &fake_image_repository.FakeImageRepository{}
	var randomList = fakeImage.GetImageEndpoint(10)
	var valueList = len(randomList)
	var valueCompare = 10

	// if len(randomList) == 10 {
	// 	t.Log("El test ha finalizado correctamente")
	// } else {
	// 	t.Error("la lista esta vacia")
	// }

	assert.Equal(t, valueList, valueCompare, "The two variables should be the same length.")

}
