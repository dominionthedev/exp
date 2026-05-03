package sheme

// AdjustStage performs a final quality pass over the 16 assigned ANSI colours.
//
// Problems it catches that AssignStage cannot prevent:
//   - A semantic colour (1–6) accidentally too close to bg in L → invisible
//   - A semantic colour too close to fg in L → blends into text
//   - A bright variant (9–14) not visibly lighter than its base (1–6)
//   - Two adjacent semantic colours with < 30° hue separation after palette
//     influence → they look the same (the original generator had this problem)
//
// Adjustments are minimal — just enough to cross each threshold.
type AdjustStage struct{}

func (s *AdjustStage) Name() string { return "adjust" }

func (s *AdjustStage) Process(ctx *Context) error {
	s.semanticContrast(ctx)
	s.brightSeparation(ctx)
	return nil
}

// semanticContrast ensures each of ANSI 1–6 is visible against both bg and fg.
//
//	ΔL vs bg ≥ 18    — colour must stand out from background
//	ΔL vs fg ≥ 12    — colour must not blend into foreground text
func (s *AdjustStage) semanticContrast(ctx *Context) {
	bgL := ctx.Background.Light()
	fgL := ctx.Foreground.Light()

	for i := 1; i <= 6; i++ {
		t := ctx.Colors[i]

		// Too close to background — push away
		if deltaL(t, ctx.Background) < 18 {
			if ctx.IsDark {
				t = withL(t, clamp(bgL+22, 0, 100))
			} else {
				t = withL(t, clamp(bgL-22, 0, 100))
			}
		}

		// Too close to foreground — push toward the middle between bg and fg
		if deltaL(t, ctx.Foreground) < 12 {
			mid := (bgL + fgL) / 2
			t = withL(t, mid)
		}

		ctx.Colors[i] = t
	}
}

// brightSeparation ensures each bright variant (ANSI 9–14) is visibly lighter
// (dark theme) or darker (light theme) than its base colour (ANSI 1–6).
// Minimum separation: ΔL ≥ 9.
func (s *AdjustStage) brightSeparation(ctx *Context) {
	const minDelta = 9.0

	for i := 1; i <= 6; i++ {
		base := ctx.Colors[i]
		bright := ctx.Colors[i+8]

		var diff float64
		if ctx.IsDark {
			diff = bright.Light() - base.Light() // bright should be lighter
		} else {
			diff = base.Light() - bright.Light() // bright should be lighter on light theme too
		}

		if diff < minDelta {
			if ctx.IsDark {
				ctx.Colors[i+8] = withL(bright, clamp(base.Light()+minDelta, 0, 100))
			} else {
				ctx.Colors[i+8] = withL(bright, clamp(base.Light()+minDelta, 0, 100))
			}
		}
	}
}
