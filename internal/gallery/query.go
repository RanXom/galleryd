package gallery

type SortField string

const (
	SortByDateTime SortField = "dateTaken"
	SortByPath     SortField = "path"
)

// SortOrder identifies the sort direction
type SortOrder string

const (
	SortAsc  SortOrder = "asc"
	SortDesc SortOrder = "desc"
)

// Query describes how photos should be selected.
//
// The zero value returns the entire gallery.
type Query struct {
	Limit  int
	Offset int

	Sort  SortField
	Order SortOrder
}
