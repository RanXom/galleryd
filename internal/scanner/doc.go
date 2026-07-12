// Package scanner discovers files from one or more filesystem roots.
//
// The scanner is intentionally simple:
//
//   - It walks configured directories.
//   - It returns discovered files.
//   - It does not read image metadata.
//   - It does not generate thumbnails.
//   - It does not sort results.
//   - It does not cache.
//
// Higher-level packages are responsible for interpreting the discovered
// files and building a photo gallery from them.
package scanner
