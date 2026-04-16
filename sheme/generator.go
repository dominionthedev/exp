package sheme

import (
	"github.com/leraniode/wondertone/core"
	"github.com/leraniode/wondertone/palette"
)

// Generator handles conversion from a Wondertone Palette to a terminal Theme.
type Generator struct {
	BasePalette *palette.Palette
}

// NewGenerator creates a new generator for the given palette.
func NewGenerator(p *palette.Palette) *Generator {
	return &Generator{BasePalette: p}
}

// GenerateTheme maps the Wondertone palette tones to terminal ANSI colors and sequences.
func (g *Generator) GenerateTheme() *Theme {
	t := &Theme{
		Name:   g.BasePalette.Name(),
		Author: g.BasePalette.Author(),
	}

	// Foreground and Background
	bg, _ := g.BasePalette.Get(g.BasePalette.Name() + " Base")
	if bg == (core.Tone{}) {
		bg, _ = g.BasePalette.Get("Base")
	}
	t.Background = bg

	fg, _ := g.BasePalette.Get(g.BasePalette.Name() + " Text")
	if fg == (core.Tone{}) {
		fg, _ = g.BasePalette.Get("Text")
	}
	t.Foreground = fg

	// Cursor defaults to Accent
	accent, _ := g.BasePalette.Get(g.BasePalette.Name() + " Accent")
	if accent == (core.Tone{}) {
		accent, _ = g.BasePalette.Get("Accent")
	}
	t.Cursor = accent

	// Map ANSI Colors (0-15)
	l := 60.0
	v := 70.0
	if accent != (core.Tone{}) {
		l = accent.Light()
		v = accent.Vibrancy()
	}

	// Index 0: Black (Background)
	t.Colors[0] = bg

	// Index 1-6: Red, Green, Yellow, Blue, Magenta, Cyan
	ansiHues := []float64{14, 142, 38, 240, 320, 196}
	for i, hue := range ansiHues {
		t.Colors[i+1] = core.New(core.Light(l), core.Vibrancy(v), core.Hue(hue))
	}

	// Index 7: White (Foreground)
	t.Colors[7] = fg

	// Bright versions (8-15)
	t.Colors[8] = bg.Lighten(15) // Bright Black
	for i := 1; i <= 6; i++ {
		t.Colors[i+8] = t.Colors[i].Lighten(10).Saturate(5)
	}
	t.Colors[15] = fg.Lighten(10) // Bright White

	return t
}
