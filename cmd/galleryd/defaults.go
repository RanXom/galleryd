package main

import (
	"log"
	"os"
	"path/filepath"
)

// defaultPicturesDir returns the user's Pictures directory.
func defaultPicturesDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, "Pictures")
}

// defaultCacheDir returns the default galleryd cache directory.
func defaultCacheDir() string {
	cache, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(cache, "galleryd")
}
