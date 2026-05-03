package sheme

import "github.com/leraniode/wondertone/core"

// ansiHues defines the canonical hue for each semantic ANSI colour slot 1–6.
// Indices map to: red, green, yellow, blue, magenta, cyan.
//
// Chosen for maximum hue separation (each slot ≥ 55° from its neighbours)
// while staying perceptually close to the universal terminal colour meaning.
var ansiHues = [6]float64{
	2,   // red     (ANSI 1) — near-pure red
	132, // green   (ANSI 2) — natural mid-green
	44,  // yellow  (ANSI 3) — warm amber-yellow
	222, // blue    (ANSI 4) — standard terminal blue
	292, // magenta (ANSI 5) — purple-magenta
	188, // cyan    (ANSI 6) — teal-cyan
}

// AssignStage maps the normalized + neutralized context tones to ANSI 0–15
// colour slots. Dark and light palettes use different L/V strategies so that
// semantic colours (red for errors, green for success, etc.) remain legible
// in both orientations.
type AssignStage struct{}

func (s *AssignStage) Name() string { return "assign" }

func (s *AssignStage) Process(ctx *Context) error {
	if ctx.IsDark {
		s.dark(ctx)
	} else {
		s.light(ctx)
	}
	return nil
}

// dark assigns colours for dark-background palettes.
//
// Strategy:
//   - ANSI 0 (black)        = background itself
//   - ANSI 1–6 (semantics)  = mid-bright colours derived from accent L/V + canonical hues
//   - ANSI 7 (white)        = foreground
//   - ANSI 8 (bright black) = background + Δ12L
//   - ANSI 9–14 (brights)   = semantic colours + Δ12L + Δ8V
//   - ANSI 15 (bright white)= foreground + Δ8L, capped at 98
func (s *AssignStage) dark(ctx *Context) {
	acL := ctx.Accent.Light()
	acV := ctx.Accent.Vibrancy()

	// Semantic colour L/V: anchored to accent but kept in a readable band.
	semL := clamp(acL*0.92, 42, 66)
	semV := clamp(acV, 55, 84)

	ctx.Colors[0] = ctx.Background
	for i, h := range ansiHues {
		ctx.Colors[i+1] = fromHLV(h, semL, semV)
	}
	ctx.Colors[7] = ctx.Foreground

	ctx.Colors[8] = withL(ctx.Background, clamp(ctx.Background.Light()+12, 0, 40))
	for i := 1; i <= 6; i++ {
		base := ctx.Colors[i]
		ctx.Colors[i+8] = withLV(base, clamp(base.Light()+12, 0, 100), clamp(base.Vibrancy()+8, 0, 100))
	}
	ctx.Colors[15] = withL(ctx.Foreground, clamp(ctx.Foreground.Light()+8, 0, 98))
}

// light assigns colours for light-background palettes.
//
// Strategy:
//   - ANSI 0 (black)        = near-pure dark (independent of palette)
//   - ANSI 1–6 (semantics)  = deeper, more saturated colours (readable on light bg)
//   - ANSI 7 (white)        = background itself
//   - ANSI 8 (bright black) = dark grey
//   - ANSI 9–14 (brights)   = semantic colours + Δ8L
//   - ANSI 15 (bright white)= near-pure white (independent of palette)
func (s *AssignStage) light(ctx *Context) {
	acL := ctx.Accent.Light()
	acV := ctx.Accent.Vibrancy()

	semL := clamp(acL*0.65, 22, 48)
	semV := clamp(acV, 58, 82)

	ctx.Colors[0] = core.New(core.Light(5), core.Vibrancy(0), core.Hue(0))
	for i, h := range ansiHues {
		ctx.Colors[i+1] = fromHLV(h, semL, semV)
	}
	ctx.Colors[7] = ctx.Background

	ctx.Colors[8] = core.New(core.Light(28), core.Vibrancy(0), core.Hue(0))
	for i := 1; i <= 6; i++ {
		base := ctx.Colors[i]
		ctx.Colors[i+8] = withL(base, clamp(base.Light()+8, 0, 100))
	}
	ctx.Colors[15] = core.New(core.Light(98), core.Vibrancy(0), core.Hue(0))
}
