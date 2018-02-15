package interfaces

import (
	"image"

	"github.com/co0p/tics/domain"
)

// StorageHandler allows us to swap in different implementations of the actual storage method
type StorageHandler interface {
	Add(hash string, image image.Image)
	Get(hash string) (image.Image, bool)
}

type ThumbnailStorageRepo struct {
	StorageHandler StorageHandler
}

func NewThumbnailStorageRepo(handler StorageHandler) *ThumbnailStorageRepo {
	return &ThumbnailStorageRepo{StorageHandler: handler}
}

func (tr *ThumbnailStorageRepo) Add(thumbnail domain.Thumbnail) {
	tr.StorageHandler.Add(thumbnail.GetHash(), thumbnail.GetImage())
}

func (tr *ThumbnailStorageRepo) Get(hash string) (domain.Thumbnail, bool) {
	img, found := tr.StorageHandler.Get(hash)
	if !found {
		return domain.Thumbnail{}, false
	}
	return domain.NewThumbnail(hash, img), true
}
