package infrastructure

import "image"

// CappedMemoryStorage is a simple in memory storage for images
type CappedMemoryStorage struct {
	thumbnails map[string]image.Image
}

// NewCappedMemoryStorage returns a new initialized CappedMemoryStorage with initial capacity
func NewCappedMemoryStorage(capacity int) *CappedMemoryStorage {
	return &CappedMemoryStorage{make(map[string]image.Image)}
}

// Add adds a new image to the storage under the key hash
func (ms *CappedMemoryStorage) Add(hash string, image image.Image) {
	ms.thumbnails[hash] = image
}

// Get returns the image and a found flag for the given hash
func (ms *CappedMemoryStorage) Get(hash string) (image.Image, bool) {
	img, found := ms.thumbnails[hash]
	if !found {
		return nil, false
	}
	return img, true
}
