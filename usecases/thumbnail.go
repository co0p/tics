// Package usecases is the bridge between the domain and the interfaces
package usecases

import "github.com/co0p/tics/domain"

import "image"

// Logger abstracts away the loggin
type Logger interface {
	Log(format string, args ...interface{})
}

// ImageResizer resizes an image based on w and h
type ImageResizer interface {
	Resize(image image.Image, w, h int) (image.Image, error)
}

// HashDecoder decodes a hash to an ImageDescription
type HashDecoder interface {
	Decode(string) (ImageDescription, error)
}

// ImageFetcher fetches an image from the given url.URL
type ImageFetcher interface {
	Fetch(path string) (image.Image, error)
}

// ImageDescription describes an image extracted from the hash
type ImageDescription struct {
	Path   string
	Width  int
	Height int
}

// ThumbnailInteractor contains all players involved in creating thumbnails
type ThumbnailInteractor struct {
	ThumbnailRepository domain.ThumbnailRepository
	ImageResizer        ImageResizer
	ImageFetcher        ImageFetcher
	HashDecoder         HashDecoder
	Logger              Logger
}

// Add adds a new thumbnail based on the hash to the thumbnail repository
func (ti *ThumbnailInteractor) Add(hash string) (image.Image, error) {
	ti.Logger.Log("Add new thumbnail with hash:%s", hash)

	if thumbnail, found := ti.ThumbnailRepository.Get(hash); found {
		ti.Logger.Log("There is a already a thumbnail with this hash; abort.\n")
		return thumbnail.GetImage(), nil
	}

	desc, err := ti.HashDecoder.Decode(hash)
	if err != nil {
		ti.Logger.Log("failed decoding: %v", err)
		return nil, err
	}

	oImage, err := ti.ImageFetcher.Fetch(desc.Path)
	if err != nil {
		ti.Logger.Log("failed fetching image: %v", err)
		return nil, err
	}

	rImage, err := ti.ImageResizer.Resize(oImage, desc.Width, desc.Height)
	if err != nil {
		ti.Logger.Log("failed resizing: %v", err)
		return nil, err
	}

	ti.ThumbnailRepository.Add(domain.NewThumbnail(hash, rImage))
	ti.Logger.Log("Adding new thumbnail with hash:%s is done!\n", hash)
	return rImage, nil
}

// Get returns an thumbnail image generated from the given hash, error otherwise
func (ti *ThumbnailInteractor) Get(hash string) (image.Image, error) {
	if thumbnail, found := ti.ThumbnailRepository.Get(hash); found {
		ti.Logger.Log("Found thumbnail for hash:%s!\n", hash)
		return thumbnail.GetImage(), nil
	}

	ti.Logger.Log("Did not find thumbnail for hash: %s, generating ...\n", hash)
	thumbnail, err := ti.Add(hash)
	if err != nil {
		ti.Logger.Log("failed adding new thumbnail: %v", err)
		return nil, err
	}

	ti.Logger.Log("Generating thumbnail for hash:%s is done!\n", hash)
	return thumbnail, nil
}
