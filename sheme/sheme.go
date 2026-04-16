package sheme

import (
	"fmt"
	"github.com/leraniode/wondertone/palette"
)

// Generate takes a Wondertone palette and returns the .theme content.
func Generate(p *palette.Palette) (string, error) {
	if p == nil {
		return "", fmt.Errorf("palette cannot be nil")
	}
	g := NewGenerator(p)
	theme := g.GenerateTheme()
	return theme.Export(), nil
}
