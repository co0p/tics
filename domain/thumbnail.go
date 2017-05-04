// Package domain contains the domain entities and domain logic necessary for
// the thumbnail image cache.
package domain

import (
	"image"
)

// ThumbnailRepository allow to add and retrieve thumbnails
type ThumbnailRepository interface {
	Add(Thumbnail)
	Get(hash string) (Thumbnail, bool)
}

// Thumbnail encapsulates the actual thumbnail image but also tracks access time
type Thumbnail struct {
	hash  string
	image image.Image
}

// GetImage returns the image data and updates the internal data accordingly
func (t *Thumbnail) GetImage() image.Image {
	return t.image
}

// GetHash returns the image hash
func (t *Thumbnail) GetHash() string {
	return t.hash
}

// NewThumbnail creates a new thumbnail tying the hash with the image together
func NewThumbnail(hash string, image image.Image) Thumbnail {
	return Thumbnail{hash: hash, image: image}
}
