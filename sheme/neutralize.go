package sheme

// NeutralizeStage fixes problematic colour conditions that cause poor
// visibility in real terminal emulators. It operates after normalization,
// so it works with already-clamped base tones.
//
// Problems addressed, in order:
//
//  1. fg/bg contrast too low    — text becomes unreadable
//  2. Accent invisible on bg    — cursor and highlights disappear
//  3. Accent vibrancy too low   — washed-out cursor / highlight
//  4. Accent vibrancy too high  — eye-burning neon cursor
//  5. Accent hue too close to fg — cursor blends into text
//
// Each rule adjusts the minimum amount necessary to cross the threshold —
// it does not try to make colours "pretty", only safe.
type NeutralizeStage struct{}

func (s *NeutralizeStage) Name() string { return "neutralize" }

func (s *NeutralizeStage) Process(ctx *Context) error {
	s.enforceTextContrast(ctx)
	s.enforceAccentVisibility(ctx)
	s.enforceAccentVibrancy(ctx)
	s.enforceAccentHueSeparation(ctx)
	ctx.Cursor = ctx.Accent
	return nil
}

// enforceTextContrast ensures the WCAG contrast ratio between fg and bg is ≥ 4.5.
// This is the most critical rule — if text is invisible, nothing else matters.
func (s *NeutralizeStage) enforceTextContrast(ctx *Context) {
	const minRatio = 4.5
	bgL := ctx.Background.Light()
	fgL := ctx.Foreground.Light()

	if contrastRatio(bgL, fgL) >= minRatio {
		return // already fine
	}

	// Push fg further away from bg until contrast is met.
	// Use binary search over the L axis: max 16 iterations is enough.
	lo, hi := fgL, float64(0)
	if ctx.IsDark {
		lo, hi = fgL, 100.0
	} else {
		lo, hi = 0.0, fgL
	}

	for i := 0; i < 16; i++ {
		mid := lerp(lo, hi, 0.5)
		if contrastRatio(bgL, mid) >= minRatio {
			hi = mid
		} else {
			lo = mid
		}
	}
	ctx.Foreground = withL(ctx.Foreground, hi)
}

// enforceAccentVisibility ensures the accent is perceptually distinct from bg (ΔL ≥ 22).
func (s *NeutralizeStage) enforceAccentVisibility(ctx *Context) {
	const minDeltaL = 22.0
	dl := deltaL(ctx.Accent, ctx.Background)
	if dl >= minDeltaL {
		return
	}

	bgL := ctx.Background.Light()
	if ctx.IsDark {
		// Push accent lighter
		ctx.Accent = withL(ctx.Accent, clamp(bgL+minDeltaL, 0, 100))
	} else {
		// Push accent darker
		ctx.Accent = withL(ctx.Accent, clamp(bgL-minDeltaL, 0, 100))
	}
}

// enforceAccentVibrancy keeps accent vibrancy in [35, 88].
// Below 35 → colours look grey and lifeless.
// Above 88 → colours become physically painful on some displays.
func (s *NeutralizeStage) enforceAccentVibrancy(ctx *Context) {
	v := ctx.Accent.Vibrancy()
	switch {
	case v < 35:
		ctx.Accent = withV(ctx.Accent, 50)
	case v > 88:
		ctx.Accent = withV(ctx.Accent, 80)
	}
}

// enforceAccentHueSeparation ensures accent hue is ≥ 30° away from foreground hue.
// When they share a hue, the cursor blends into text — invisible in practice.
func (s *NeutralizeStage) enforceAccentHueSeparation(ctx *Context) {
	const minDeltaH = 30.0
	if deltaH(ctx.Accent, ctx.Foreground) >= minDeltaH {
		return
	}
	// Rotate accent hue by 60° — enough separation without fighting the palette.
	newHue := ctx.Accent.Hue() + 60
	if newHue >= 360 {
		newHue -= 360
	}
	ctx.Accent = fromHLV(newHue, ctx.Accent.Light(), ctx.Accent.Vibrancy())
}
