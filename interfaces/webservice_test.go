package interfaces_test

import (
	"errors"
	"image"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/co0p/tics/interfaces"
)

func TestGetThumbnailToReturnHttpErrorOnMissingQueryParameter(t *testing.T) {

	mock := MockThumbnailInteractor{}
	mock.GetFn = func(hash string) (image.Image, error) {
		return nil, nil
	}
	webserviceHandler := interfaces.WebserviceHandler{
		ThumbnailInteractor: &mock,
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(webserviceHandler.GetThumbnail)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	if len(rr.Body.String()) == 0 {
		t.Errorf("Handler should return error message")
	}
}

func TestGetThumbnailToReturnInternalErrorOnRetrivalError(t *testing.T) {
	expectedErrorMessage := "error message"
	mock := MockThumbnailInteractor{}
	mock.GetFn = func(hash string) (image.Image, error) {
		return nil, errors.New(expectedErrorMessage)
	}
	webserviceHandler := interfaces.WebserviceHandler{
		ThumbnailInteractor: &mock,
	}

	req, _ := http.NewRequest("GET", "/?i=2323", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(webserviceHandler.GetThumbnail)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	msg := strings.TrimRight(rr.Body.String(), "\n")
	if msg != expectedErrorMessage {
		t.Errorf("Expected Error message %v, but got %v", expectedErrorMessage, msg)
	}
}

func TestGetThumbnailToReturnInternalErrorOnEncodingError(t *testing.T) {

}

func TestGetThumbnailToSetAndReturnCorrectContentTypeAndImage(t *testing.T) {

}

type MockThumbnailInteractor struct {
	GetFn       func(hash string) (image.Image, error)
	GetFnCalled bool
}

func (m *MockThumbnailInteractor) Get(hash string) (image.Image, error) {
	m.GetFnCalled = true
	return m.GetFn(hash)
}
