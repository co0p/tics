package main

import (
	"log"
	"net/http"
	"os"

	"github.com/co0p/tics/infrastructure"
	"github.com/co0p/tics/interfaces"
	"github.com/co0p/tics/usecases"
)

func main() {

	// we use a memory storage service
	storage := infrastructure.NewCappedMemoryStorage(10)
	thumbnailRepo := interfaces.NewThumbnailStorageRepo(storage)

	// wireing all the components together
	interactor := new(usecases.ThumbnailInteractor)
	interactor.ThumbnailRepository = thumbnailRepo
	interactor.ImageResizer = infrastructure.MNResizer{}
	interactor.ImageFetcher = infrastructure.ImageFetcher{}
	interactor.HashDecoder = infrastructure.Base64Decoder{}
	interactor.Logger = infrastructure.ConsoleLogger{}

	// tell the webserver to use our newly configured
	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.ThumbnailInteractor = interactor

	// start handling some requests
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		webserviceHandler.GetThumbnail(res, req)
	})
	port := os.Getenv("TICS_PORT")
	if len(port) == 0 {
		port = "8080"
	}
	interactor.Logger.Log("starting server on port %v ...\n", port)
	if http.ListenAndServe(":"+port, nil) != nil {
		log.Fatalln("failed to start server")
	}
}
