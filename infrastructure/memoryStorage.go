package infrastructure

import "image"

// MemoryStorage is a simple in memory storage for images
type MemoryStorage struct {
	thumbnails map[string]image.Image
}

// NewMemoryStorage returns a new initialized MemoryStorage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{make(map[string]image.Image)}
}

// Add add a new image to the storage under the key haash
func (ms *MemoryStorage) Add(hash string, image image.Image) {
	ms.thumbnails[hash] = image
}

// Get returns the image and a found flag for the given hash
func (ms *MemoryStorage) Get(hash string) (image.Image, bool) {
	img, found := ms.thumbnails[hash]
	if !found {
		return nil, false
	}
	return img, true
}
