package infrastructure

import (
	"errors"
	"image"

	"github.com/nfnt/resize"
)

// MNResizer is an image resizer using the Mitchell-Netravali algorithm
type MNResizer struct{}

func (r MNResizer) Resize(img image.Image, w, h int) (image.Image, error) {

	if w <= 0 && h <= 0 {
		return nil, errors.New("width and height should be positive")
	}

	size := img.Bounds().Size()
	if size.X < w || size.Y < h {
		return img, nil
	}

	return resize.Resize(uint(w), uint(h), img, resize.MitchellNetravali), nil
}
