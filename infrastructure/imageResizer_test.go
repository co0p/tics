package infrastructure

import (
	"image"
	"testing"
)

func Test_Resize_should_not_resize_smaller_image(t *testing.T) {

	testImages := []image.Image{
		image.NewRGBA(image.Rect(0, 0, 10, 10)),
		image.NewRGBA(image.Rect(0, 0, 100, 100)),
		image.NewRGBA(image.Rect(0, 0, 100, 200)),
	}

	resizer := MNResizer{}
	for _, tt := range testImages {
		actual, err := resizer.Resize(tt, 100, 200)

		if err != nil {
			t.Error("Expected no error, got error:", err.Error())
		}

		actualSize := actual.Bounds().Size()
		expectedSize := tt.Bounds().Size()

		if actualSize.X != expectedSize.X || actualSize.Y != expectedSize.Y {
			t.Errorf("Expected size=%q, got: %q", expectedSize, actualSize)
		}
	}
}
func Test_Resize_should_resize_bigger_image_to_target_w_and_h(t *testing.T) {

	testImages := []image.Image{
		image.NewRGBA(image.Rect(0, 0, 100, 200)),
		image.NewRGBA(image.Rect(0, 0, 200, 200)),
		image.NewRGBA(image.Rect(0, 0, 300, 200)),
		image.NewRGBA(image.Rect(0, 0, 200, 300)),
		image.NewRGBA(image.Rect(0, 0, 300, 300)),
	}

	resizer := MNResizer{}
	for _, tt := range testImages {
		actual, err := resizer.Resize(tt, 100, 200)

		if err != nil {
			t.Error("Expected no error, got error:", err.Error())
		}

		size := actual.Bounds().Size()
		if size.X != 100 || size.Y != 200 {
			t.Errorf("Expected size=(100,200), got: %q", size)
		}
	}
}
