package sheme

import (
	"strings"
)

// TestHelper helps in testing .theme files.
type TestHelper struct{}

// IsValidTheme checks if the content looks like a valid .theme file.
func (h *TestHelper) IsValidTheme(content string) bool {
	return strings.Contains(content, "# Color Tokens") &&
		strings.Contains(content, "export COLOR_FG") &&
		strings.Contains(content, "# OSC Sequences") &&
		strings.Contains(content, "printf \"\\033]10;")
}
