package interfaces_test

import (
	"image"
	"testing"

	"github.com/co0p/tics/domain"
	"github.com/co0p/tics/interfaces"
)

func TestAddToDelegateToStorage(t *testing.T) {
	mockStorage := MockStorage{
		AddFn: func(hash string, image image.Image) {},
		GetFn: func(hash string) (image.Image, bool) { return nil, false },
	}
	thumbnail := domain.Thumbnail{}
	repo := interfaces.NewThumbnailStorageRepo(&mockStorage)
	repo.Add(thumbnail)

	if !mockStorage.AddFnCalled {
		t.Errorf("Add of storage handler should have been called")
	}
}

func TestGetToReturnImage(t *testing.T) {
	imageHash := "123123123"
	img := image.NewRGBA(image.Rect(0, 0, 0, 0))
	mockStorage := MockStorage{
		GetFn: func(hash string) (image.Image, bool) {
			return img, true
		},
	}
	repo := interfaces.NewThumbnailStorageRepo(&mockStorage)
	thumbnail, found := repo.Get(imageHash)

	if !mockStorage.GetFnCalled {
		t.Errorf("Get of storage handler should have been called")
	}

	if thumbnail.GetImage() != img {
		t.Errorf("expected img of thumbnail to be same as input")
	}
	if thumbnail.GetHash() != imageHash {
		t.Errorf("expected getHash of thumbnail to be 'someHash', got %v", thumbnail.GetHash())
	}

	if !found {
		t.Errorf("expected found to be true, got %v", found)
	}
}

func TestGetToReturnNotFoundIfImageWasNotFound(t *testing.T) {
	mockStorage := MockStorage{
		GetFn: func(hash string) (image.Image, bool) {
			return nil, false
		},
	}
	repo := interfaces.NewThumbnailStorageRepo(&mockStorage)
	_, found := repo.Get("imageHash")

	if !mockStorage.GetFnCalled {
		t.Errorf("Get of storage handler should have been called")
	}

	if found {
		t.Errorf("expected found to be false, got %t", found)
	}
}

type MockStorage struct {
	AddFn       func(hash string, image image.Image)
	AddFnCalled bool
	GetFn       func(hash string) (image.Image, bool)
	GetFnCalled bool
}

func (s *MockStorage) Add(hash string, image image.Image) {
	s.AddFnCalled = true
	s.AddFn(hash, image)
}

func (s *MockStorage) Get(hash string) (image.Image, bool) {
	s.GetFnCalled = true
	return s.GetFn(hash)
}
