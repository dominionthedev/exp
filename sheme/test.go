package sheme

import (
	"fmt"
	"strings"
)

// TestHelper provides assertions for rendered theme content.
// Used in tests to validate the structural correctness of Generate() output.
type TestHelper struct{}

// IsValidTheme checks that the rendered string has all required sections.
func (h *TestHelper) IsValidTheme(content string) bool {
	required := []string{
		"Color tokens",
		`export COLOR_FG="`,
		`export COLOR_BG="`,
		`export COLOR_CURSOR="`,
		"Terminal sequences",
		`printf "\033]10;`,
		`printf "\033]11;`,
		`printf "\033]12;`,
	}
	for _, r := range required {
		if !strings.Contains(content, r) {
			return false
		}
	}
	return true
}

// HasAllANSI checks that all 16 ANSI colour slots (COLOR_0..COLOR_15) are exported.
func (h *TestHelper) HasAllANSI(content string) bool {
	for i := range 16 {
		if !strings.Contains(content, fmt.Sprintf(`export COLOR_%d="`, i)) {
			return false
		}
	}
	return true
}

// HasAllOSC checks that OSC sequences for all 16 ANSI slots are present.
func (h *TestHelper) HasAllOSC(content string) bool {
	for i := range 16 {
		if !strings.Contains(content, fmt.Sprintf(`\033]4;%d;`, i)) {
			return false
		}
	}
	return true
}
