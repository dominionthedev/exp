package sheme

// ClassifyStage analyses the extracted base tones and sets palette-wide
// classification flags that later stages use to make informed decisions.
//
// Classification results:
//   IsDark         — background lightness ≤ 55 (terminal dark-mode threshold)
//   IsHighContrast — fg/bg ΔL > 70 (very high contrast pair)
type ClassifyStage struct{}

func (s *ClassifyStage) Name() string { return "classify" }

func (s *ClassifyStage) Process(ctx *Context) error {
	bgL := ctx.Background.Light()
	fgL := ctx.Foreground.Light()

	ctx.IsDark = bgL <= 55

	dl := deltaL(ctx.Background, ctx.Foreground)
	ctx.IsHighContrast = dl > 70

	// Sanity: if the palette has light bg but darker fg set correctly,
	// and the user's palette happens to have fg lighter than bg (inverted),
	// swap them so downstream stages always operate on bg-dark / fg-light
	// for dark themes and the inverse for light themes.
	if ctx.IsDark && fgL < bgL {
		ctx.Background, ctx.Foreground = ctx.Foreground, ctx.Background
	}
	if !ctx.IsDark && fgL > bgL {
		ctx.Background, ctx.Foreground = ctx.Foreground, ctx.Background
	}

	return nil
}
