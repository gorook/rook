package main

import (
	"os"
	"testing"
)

// just cover main()
func TestMainFunc(t *testing.T) {
	old := os.Stderr
	_, w, _ := os.Pipe()
	os.Stderr = w
	main()
	os.Stderr = old
}
