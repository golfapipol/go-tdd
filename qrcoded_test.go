package main

import (
	"errors"
	"testing"
)

type ErrorWriter struct{}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error")
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "555-2368")

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Generated QRCode is not a PNG: %s", err)
	}
}
