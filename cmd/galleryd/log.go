package main

import (
	"fmt"
	"strings"
)

func listenURL(addr string) string {
	switch {
	case strings.HasPrefix(addr, ":"):
		return fmt.Sprintf("http://localhost%s", addr)
	default:
		return fmt.Sprintf("http://%s", addr)
	}
}
