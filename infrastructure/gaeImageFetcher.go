package infrastructure

import (
	"image"
	"image/png"
	"net/http"
)

// Google App engine ImageFetcher implements the ImageFetcher usecases interface
type GaeImageFetcher struct {
	Client *http.Client
}

// Fetch fetches and returns the image found unter url, error otherwise
func (i GaeImageFetcher) Fetch(path string) (image.Image, error) {
	resp, err := i.Client.Get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return png.Decode(resp.Body)
}
