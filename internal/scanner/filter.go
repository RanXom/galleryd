package scanner

import "strings"

func shouldSkipDir(name string) bool {
	return strings.HasPrefix(name, ".") && name != "."
}

func shouldSkipFile(name string) bool {
	return strings.HasPrefix(name, ".")
}
