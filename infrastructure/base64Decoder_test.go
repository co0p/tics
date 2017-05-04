package infrastructure

import (
	"testing"

	"github.com/co0p/tics/usecases"
)

func TestBase64Decoder_should_decode_and_parse_simple_hash(t *testing.T) {
	in := "aHR0cDovL3d3dy5pbWcuZGUvaW1hZ2UucG5nP3c9MTAmaD0yMA=="
	expected := usecases.ImageDescription{
		Path:   "http://www.img.de/image.png?w=10&h=20",
		Width:  10,
		Height: 20,
	}

	decoder := Base64Decoder{}
	actual, err := decoder.Decode(in)

	if err != nil {
		t.Errorf("Decode(%s) => %q unexpected error: %s", in, actual, err)
	}

	if !isEqual(actual, expected) {
		t.Errorf("Decode(%s) => %q, want %q", in, actual, expected)
	}

}

func TestBase64Decoder_should_decode_and_parse_complex_hash(t *testing.T) {
	in := "aHR0cHM6Ly91cGxvYWQud2lraW1lZGlhLm9yZy93aWtpcGVkaWEvY29tbW9ucy90aHVtYi80LzQyL1NhbXBsZS1pbWFnZS5zdmcvMTI4MHB4LVNhbXBsZS1pbWFnZS5zdmcucG5nP3c9MTAwJmg9MjAw"
	expected := usecases.ImageDescription{
		Path:   "https://upload.wikimedia.org/wikipedia/commons/thumb/4/42/Sample-image.svg/1280px-Sample-image.svg.png?w=100&h=200",
		Width:  100,
		Height: 200,
	}

	decoder := Base64Decoder{}
	actual, err := decoder.Decode(in)

	if err != nil {
		t.Errorf("Decode(%s) => %q unexpected error: %s", in, actual, err)
	}

	if !isEqual(actual, expected) {
		t.Errorf("Decode(%s) => %q, want %q", in, actual, expected)
	}
}

func isEqual(a usecases.ImageDescription, b usecases.ImageDescription) bool {
	return (a.Height == b.Height) && (a.Width == b.Width) && (a.Path == b.Path)
}
