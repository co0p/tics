package infrastructure

import (
	b64 "encoding/base64"
	"errors"
	"net/url"

	"strconv"

	"github.com/co0p/tics/usecases"
)

// Base64Decoder is a simple base64 decoder
type Base64Decoder struct{}

func decode(input string) (string, error) {
	decoded, err := b64.StdEncoding.DecodeString(input)

	if err != nil {
		return "", errors.New("failed to decode string using base64:" + err.Error())
	}

	return string(decoded), nil
}

// Decode decodes the given hash and returns an imageDescription, error otherwise
func (d Base64Decoder) Decode(hash string) (usecases.ImageDescription, error) {
	decoded, err := decode(hash)
	if err != nil {
		return usecases.ImageDescription{}, err
	}

	url, err := url.Parse(decoded)
	if err != nil {
		return usecases.ImageDescription{}, errors.New("Failed to parse decoded string into valid url")
	}

	queryValues := url.Query()

	width, err := strconv.Atoi(queryValues.Get("w"))
	if err != nil {
		return usecases.ImageDescription{}, errors.New("decoded string does not contains valid width at query w")
	}

	height, err := strconv.Atoi(queryValues.Get("h"))
	if err != nil {
		return usecases.ImageDescription{}, errors.New("decoded string does not contains valid height at query h")
	}

	return usecases.ImageDescription{
		Path:   decoded, // for the image path, we use the complete path including any query parameters
		Width:  width,
		Height: height,
	}, nil
}
