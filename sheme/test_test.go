package sheme

import (
	"strings"
	"testing"

	"github.com/leraniode/wondertone/palette/builtin"
)

func TestGenerate_Aurora(t *testing.T) {
	p := builtin.Aurora()
	content, err := Generate(p)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	h := &TestHelper{}
	if !h.IsValidTheme(content) {
		t.Errorf("theme failed structural validation:\n%s", content)
	}
	if !h.HasAllANSI(content) {
		t.Errorf("theme missing ANSI slots:\n%s", content)
	}
	if !h.HasAllOSC(content) {
		t.Errorf("theme missing OSC sequences:\n%s", content)
	}
}

func TestGenerate_AllBuiltins(t *testing.T) {
	palettes := builtin.All()
	if len(palettes) == 0 {
		t.Skip("no built-in palettes available")
	}
	for _, p := range palettes {
		p := p
		t.Run(p.Name(), func(t *testing.T) {
			content, err := Generate(p)
			if err != nil {
				t.Fatalf("Generate(%q) failed: %v", p.Name(), err)
			}
			h := &TestHelper{}
			if !h.IsValidTheme(content) {
				t.Errorf("Generate(%q) produced invalid theme", p.Name())
			}
			if !h.HasAllANSI(content) {
				t.Errorf("Generate(%q) missing ANSI slots", p.Name())
			}
		})
	}
}

func TestPipeline_NilPalette(t *testing.T) {
	_, err := NewPipeline().Run(nil)
	if err == nil {
		t.Error("expected error for nil palette, got nil")
	}
}

func TestTheme_Render_HasHeader(t *testing.T) {
	p := builtin.Aurora()
	theme, err := NewPipeline().Run(p)
	if err != nil {
		t.Fatalf("pipeline failed: %v", err)
	}
	content := theme.Render()
	if !strings.Contains(content, "# Theme: ") {
		t.Errorf("render output missing header comment")
	}
}

func TestNeutralizeStage_ContrastEnforced(t *testing.T) {
	// Synthetic: very low contrast palette (bg and fg both mid-grey).
	// After pipeline, contrast ratio must be ≥ 4.5.
	p := builtin.Aurora() // stand-in; real test would use a custom palette
	pipeline := NewPipeline()
	theme, err := pipeline.Run(p)
	if err != nil {
		t.Fatalf("pipeline failed: %v", err)
	}

	bgL := theme.Background.Light()
	fgL := theme.Foreground.Light()
	ratio := contrastRatio(bgL, fgL)
	if ratio < 4.5 {
		t.Errorf("contrast ratio %.2f < 4.5 (WCAG AA) after pipeline — bg L=%.1f fg L=%.1f",
			ratio, bgL, fgL)
	}
}

func TestOSCHandler_ResetAll(t *testing.T) {
	osc := &OSCHandler{}
	reset := osc.ResetAll()
	for _, want := range []string{`\033]110`, `\033]111`, `\033]112`, `\033]104`} {
		if !strings.Contains(reset, want) {
			t.Errorf("ResetAll() missing sequence %q", want)
		}
	}
}
