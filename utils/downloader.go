package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Download downloads file at given URL to given directory with given name
// returns location of the file and error if occur
func Download(url, dir, name string) (location string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("cannot download %s: %w", url, err)
	}
	defer resp.Body.Close()

	// Create a empty file
	location = path.Join(dir, name)
	f, err := os.Create(location)
	if err != nil {
		return "", fmt.Errorf("cannot create %s: %w", location, err)
	}
	defer f.Close()

	// Write the bytes to the file
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot write %s: %w", location, err)
	}

	return location, nil
}
