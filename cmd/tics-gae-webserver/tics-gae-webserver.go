package tics

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/co0p/tics/infrastructure"
	"github.com/co0p/tics/interfaces"
	"github.com/co0p/tics/usecases"
)

func init() {

	// we use a memory storage service
	thumbnailRepo := interfaces.NewThumbnailStorageRepo(infrastructure.NewMemoryStorage())

	// wireing all the components together
	thumbnailInteractor := new(usecases.ThumbnailInteractor)
	thumbnailInteractor.ThumbnailRepository = thumbnailRepo
	thumbnailInteractor.ImageResizer = infrastructure.MNResizer{}
	thumbnailInteractor.HashDecoder = infrastructure.Base64Decoder{}
	thumbnailInteractor.Logger = infrastructure.ConsoleLogger{}

	// tell the webserver to use our newly configured thumbnailInteractor
	webserviceHandler := interfaces.WebserviceHandler{}
	webserviceHandler.ThumbnailInteractor = thumbnailInteractor

	// start handling some requests and configure the things that require a context
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		ctx := appengine.NewContext(req)
		thumbnailInteractor.Logger = infrastructure.GaeLogger{Ctx: ctx}
		thumbnailInteractor.ImageFetcher = infrastructure.GaeImageFetcher{Client: urlfetch.Client(ctx)}
		webserviceHandler.GetThumbnail(res, req)
	})
}
