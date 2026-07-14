// Package watcher watches configured directories for filesystem changes.
//
// The package is intentionally unaware of galleries, thumbnails and HTTP.
//
// Its sole responsibility is detecting filesystem changes and notifying
// the caller through a callback.
package watcher
