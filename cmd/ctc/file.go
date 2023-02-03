package main

import (
	"io"
	"log"
	"os"
)

// Read the content of a text file and return it as a string.
func Read(fname string) (string, error) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer f.Close()
	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(content), nil
}
