// Package testfixtures holds shared fixture helpers
package testfixtures

import "path/filepath"

func Canon40D() string {
	return filepath.Join("..", "testfixtures", "Canon_40D.jpg")
}

func NikonD70() string {
	return filepath.Join("..", "testfixtures", "Nikon_D70.jpg")
}

func NoEXIF() string {
	return filepath.Join("..", "testfixtures", "no-exif.png")
}
