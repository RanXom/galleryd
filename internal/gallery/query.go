package gallery

// Query describes how photos should be selected.
//
// The zero value returns the entire gallery.
type Query struct {
	Limit  int
	Offset int
}
