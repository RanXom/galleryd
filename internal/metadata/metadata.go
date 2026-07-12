package metadata

import "time"

// Metadata describes information extracted from an image file
type Metadata struct {
	Width       int
	Height      int
	DateTaken   time.Time
	Orientation int
}
