package image_domain

import (
	"reflect"
	"testing"
)

func TestImage(t *testing.T) {
	want := &Image{
		ID:    10,
		Url:   "www.mercadolibre.com",
		Autor: "camilo",
		Alt:   "garces",
		Size: &Size{
			Width:  10,
			Height: 20,
		},
	}

	got := &Image{
		ID:    10,
		Url:   "www.mercadolibre.com",
		Autor: "camilo",
		Alt:   "garces",
		Size: &Size{
			Width:  10,
			Height: 20,
		},
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("se esperaba: %v, se obtuvo %v", want, got)
	}
}
