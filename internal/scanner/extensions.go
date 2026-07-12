package scanner

import (
	"path/filepath"
	"strings"
)

// supportedExtensions contains every image extension recognised
// by the scanner.
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
