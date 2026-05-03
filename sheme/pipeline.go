package sheme

import (
	"fmt"

	"github.com/leraniode/wondertone/core"
	"github.com/leraniode/wondertone/palette"
)

// Stage is a single step in the palette → theme conversion pipeline.
// Each stage reads from Context, modifies it, and returns any error.
type Stage interface {
	Name() string
	Process(ctx *Context) error
}

// Context carries all colour state through the pipeline.
// Stages read from it, mutate it, and pass it forward.
type Context struct {
	// Input
	Palette *palette.Palette

	// Extracted base tones (set by ExtractStage)
	Background core.Tone
	Foreground core.Tone
	Cursor     core.Tone
	Accent     core.Tone

	// Palette classification (set by ClassifyStage)
	IsDark         bool
	IsHighContrast bool

	// ANSI 0–15 slot assignments (set by AssignStage, refined by AdjustStage)
	Colors [16]core.Tone
}

// Theme builds a Theme from the completed context.
func (c *Context) Theme() *Theme {
	return &Theme{
		Name:       c.Palette.Name(),
		Author:     c.Palette.Author(),
		Background: c.Background,
		Foreground: c.Foreground,
		Cursor:     c.Cursor,
		Colors:     c.Colors,
	}
}

// Pipeline runs a sequence of stages converting a palette to a Theme.
type Pipeline struct {
	stages []Stage
}

// NewPipeline returns a Pipeline with the full default stage sequence:
//
//	Extract → Classify → Normalize → Neutralize → Assign → Adjust
func NewPipeline() *Pipeline {
	return &Pipeline{
		stages: []Stage{
			&ExtractStage{},
			&ClassifyStage{},
			&NormalizeStage{},
			&NeutralizeStage{},
			&AssignStage{},
			&AdjustStage{},
		},
	}
}

// Run executes all pipeline stages in order and returns the generated Theme.
// If any stage fails, the pipeline halts and returns a wrapped error naming the stage.
func (p *Pipeline) Run(pal *palette.Palette) (*Theme, error) {
	if pal == nil {
		return nil, fmt.Errorf("palette cannot be nil")
	}
	ctx := &Context{Palette: pal}
	for _, s := range p.stages {
		if err := s.Process(ctx); err != nil {
			return nil, fmt.Errorf("[%s] %w", s.Name(), err)
		}
	}
	return ctx.Theme(), nil
}
