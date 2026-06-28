// Package sheme is the Shell Theme Engine.
//
// It converts palettes into a shell-sourceable theme string
// through a staged colour pipeline:
//
//	Extract → Classify → Normalize → Neutralize → Assign → Adjust
//
// Each stage has a single responsibility. The Neutralize stage is the key
// component that fixes colour problems (low contrast, washed-out colours,
// invisible cursor) before ANSI slots are assigned.
//
// Basic usage:
//
//	content, err := sheme.Generate(palette)
//
// Custom pipeline:
//
//	theme, err := sheme.NewPipeline().Run(palette)
//	content := theme.Render()
package sheme

import "github.com/leraniode/wondertone/palette"

// Generate converts a Wondertone palette to a shell-sourceable theme string
// using the default pipeline.
//
// The returned string is pure text — what you do with it (write to a file,
// source it, print it) is entirely up to the caller.
func Generate(p *palette.Palette) (string, error) {
	theme, err := NewPipeline().Run(p)
	if err != nil {
		return "", err
	}
	return theme.Render(), nil
}
