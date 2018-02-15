package infrastructure

import (
	"image"
	"testing"
)

func Test_Add_should_return_image_put_under_key(t *testing.T) {

	data := map[string]image.Image{
		"Hash1": image.NewRGBA(image.Rect(0, 0, 10, 10)),
		"Hash2": image.NewRGBA(image.Rect(0, 0, 20, 20)),
		"Hash3": image.NewRGBA(image.Rect(0, 0, 20, 20)),
	}

	storage := NewCappedMemoryStorage(10)

	for k, expectedItem := range data {
		storage.Add(k, expectedItem)
		actualItem, _ := storage.Get(k)

		if expectedItem != actualItem {
			t.Errorf("Expected items to be equal, %v != %v", expectedItem, actualItem)
		}
	}
}

func Test_Get_should_return_the_same_image_for_same_key(t *testing.T) {

	expectedImage := image.NewRGBA(image.Rect(0, 0, 10, 10))
	data := map[string]image.Image{
		"Hash1": expectedImage,
		"Hash2": image.NewRGBA(image.Rect(0, 0, 20, 20)),
		"Hash3": image.NewRGBA(image.Rect(0, 0, 20, 20)),
	}

	storage := NewCappedMemoryStorage(10)
	for k, v := range data {
		storage.Add(k, v)
	}

	images := []image.Image{}
	for i := 0; i < 5; i++ {
		image, _ := storage.Get("Hash1")
		images = append(images, image)
	}

	for _, actualImage := range images {

		if actualImage != expectedImage {
			t.Errorf("Expected images to be equal, %v != %v", expectedImage, actualImage)
		}
	}
}

func TestCappedMemoryStorage_should_return_false_if_no_image_was_found(t *testing.T) {
	storage := NewCappedMemoryStorage(10)
	storage.Add("Hash", image.NewRGBA(image.Rect(0, 0, 20, 20)))

	nonExistingKey := "DIFFERENT_HASH"
	thumbnail, found := storage.Get(nonExistingKey)
	if found {
		t.Errorf("expected Get(%s) to return found=false, got %t", nonExistingKey, found)
	}
	if thumbnail != nil {
		t.Errorf("expected Get(%s) to return nil, got %t", nonExistingKey, thumbnail)
	}
}
