package gallery

import (
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"

	"github.com/RanXom/galleryd/internal/scanner"
)

// generateID returns a deterministic identifier for a photo.
//
// IDs are generated from the cleaned root and relative path to avoid
// collisions when multiple roots contain the same directory layout.
func generateID(file scanner.File) string {
	canonical := filepath.Join(file.Root, file.RelativePath)

	sum := sha256.Sum256([]byte(canonical))

	return hex.EncodeToString(sum[:])
}
