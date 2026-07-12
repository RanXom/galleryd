package scanner

// Config configures a scanner.
//
// Zero values should produce sensible defaults where possible.
type Config struct {
	// Roots contains the filesystem roots that will be scanned.
	Roots []string
}
