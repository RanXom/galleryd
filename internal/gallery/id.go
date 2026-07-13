package gallery

import (
	"crypto/sha256"
	"encoding/hex"
)

func generateID(relativePath string) string {
	sum := sha256.Sum256([]byte(relativePath))
	return hex.EncodeToString(sum[:])
}
