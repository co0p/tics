package infrastructure

import (
	"image"
	"image/png"
	"net/http"
)

// ImageFetcher implements the ImageFetcher usecases interface
type ImageFetcher struct{}

// Fetch fetches and returns the image found unter url, error otherwise
func (i ImageFetcher) Fetch(path string) (image.Image, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return png.Decode(resp.Body)
}
