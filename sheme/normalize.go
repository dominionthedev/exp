package sheme

// NormalizeStage clamps Background, Foreground, and Accent into ranges
// that are physically readable in a terminal emulator.
//
// Why this matters: a palette designed for a web UI might have a background
// at L=40 (charcoal) and a foreground at L=60 (light grey) — both fine
// on a monitor but producing a muddy, low-contrast terminal theme.
//
// Ranges chosen from empirical terminal testing:
//
//	Dark  bg: L  0–22   V  0–20   (very dark, low chroma)
//	Dark  fg: L 78–96   V  0–15   (near-white, low chroma)
//	Dark  ac: L 45–72   V 45–90   (vivid, readable mid-tone)
//
//	Light bg: L 88–98   V  0–10   (near-white, very low chroma)
//	Light fg: L  8–35   V  0–15   (near-black, low chroma)
//	Light ac: L 28–55   V 45–90   (vivid, readable mid-tone)
type NormalizeStage struct{}

func (s *NormalizeStage) Name() string { return "normalize" }

func (s *NormalizeStage) Process(ctx *Context) error {
	if ctx.IsDark {
		ctx.Background = clampTone(ctx.Background, 0, 22, 0, 20)
		ctx.Foreground = clampTone(ctx.Foreground, 78, 96, 0, 15)
		ctx.Accent = clampTone(ctx.Accent, 45, 72, 45, 90)
	} else {
		ctx.Background = clampTone(ctx.Background, 88, 98, 0, 10)
		ctx.Foreground = clampTone(ctx.Foreground, 8, 35, 0, 15)
		ctx.Accent = clampTone(ctx.Accent, 28, 55, 45, 90)
	}

	// Cursor follows accent after normalization.
	ctx.Cursor = ctx.Accent
	return nil
}
