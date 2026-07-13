// Package metadata extracts information from image files.
//
// The package is responsible for reading metadata that already exists
// inside image files, such as:
//
//   - image dimensions
//   - EXIF timestamps
//   - EXIF orientation
//
// It does not generate thumbnails, assign IDs, cache results or sort
// images. Those responsibilities belong to higher-level packages.
package metadata
