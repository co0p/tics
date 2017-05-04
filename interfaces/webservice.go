package interfaces

import (
	"bytes"
	"image"
	"image/jpeg"
	"net/http"
)

type ThumbnailInteractor interface {
	Get(hash string) (image.Image, error)
}

// WebserviceHandler needs a ThumbnailInteractor to get the thumbnail generation going
type WebserviceHandler struct {
	ThumbnailInteractor ThumbnailInteractor
}

//GetThumbnail processes the incoming request and returns a resized image, the error otherwise
func (handler WebserviceHandler) GetThumbnail(w http.ResponseWriter, req *http.Request) {

	hash := req.URL.Query().Get("i")
	if len(hash) == 0 {
		http.Error(w, "Missing i query parameter", http.StatusBadRequest)
		return
	}

	image, err := handler.ThumbnailInteractor.Get(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, image, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", http.DetectContentType(buf.Bytes()))
	w.Write(buf.Bytes())
}
