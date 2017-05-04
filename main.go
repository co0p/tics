package main

import (
	"net/http"

	"github.com/co0p/tics/infrastructure"
	"github.com/co0p/tics/interfaces"
	"github.com/co0p/tics/usecases"
)

func main() {

	// we use a memory storage service
	thumbnailRepo := interfaces.NewThumbnailStorageRepo(infrastructure.NewMemoryStorage())

	// wireing all the components together
	thumbnailInteractor := new(usecases.ThumbnailInteractor)
	thumbnailInteractor.ThumbnailRepository = thumbnailRepo
	thumbnailInteractor.ImageResizer = infrastructure.MNResizer{}
	thumbnailInteractor.ImageFetcher = infrastructure.ImageFetcher{}
	thumbnailInteractor.HashDecoder = infrastructure.Base64Decoder{}
	thumbnailInteractor.Logger = infrastructure.ConsoleLogger{}

	// tell the webserver to use our newly configured thumbnailInteractor
	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.ThumbnailInteractor = thumbnailInteractor

	// start handling some requests
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.GetThumbnail(res, req)
	})
	thumbnailInteractor.Logger.Log("starting server ...")
	http.ListenAndServe(":8080", nil)
}
