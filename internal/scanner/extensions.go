package scanner

import (
	"path/filepath"
	"strings"
)

var supportedExtensions = map[string]struct{}{
	".jpg":  {},
	".jpeg": {},
	".png":  {},
	".gif":  {},
	".webp": {},
}

func isImage(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	_, ok := supportedExtensions[ext]
	return ok
}
