package scanner

import "strings"

// shouldSkipDir reports whether a directory should be skipped.
func shouldSkipDir(name string) bool {
	return strings.HasPrefix(name, ".") && name != "."
}

// shouldSkipFile reports whether a file should be skipped.
func shouldSkipFile(name string) bool {
	return strings.HasPrefix(name, ".")
}
