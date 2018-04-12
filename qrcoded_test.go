package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "555-2368", Version(1))

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}

func TestVersionDeterminesSize(t *testing.T) {
	table := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}
	for _, test := range table {
		buffer := new(bytes.Buffer)
		GenerateQRCode(buffer, "555-2368", Version(test.version))

		img, _ := png.Decode(buffer)
		if width := img.Bounds().Dx(); width != 21 {
			t.Errorf("Version %2d, expected %3d but go %3d", test.version, test.expected, width)
		}
	}

}
