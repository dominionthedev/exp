package sheme

import "github.com/leraniode/wondertone/core"

// ExtractStage pulls Background, Foreground, and Accent from the palette.
//
// Lookup order for each slot:
//  1. "<PaletteName> <Slot>"  (qualified name — most specific)
//  2. "<Slot>"                (bare name — Wondertone convention fallback)
//  3. Hardcoded default       (if palette simply does not define it)
//
// Missing tones never propagate as zero values — every slot is guaranteed
// to hold a usable tone after this stage.
type ExtractStage struct{}

func (s *ExtractStage) Name() string { return "extract" }

func (s *ExtractStage) Process(ctx *Context) error {
	n := ctx.Palette.Name()

	ctx.Background = firstTone(ctx, n+" Base", "Base")
	ctx.Foreground = firstTone(ctx, n+" Text", "Text")
	ctx.Accent = firstTone(ctx, n+" Accent", "Accent")

	// Fallback defaults — sensible dark-neutral values so downstream
	// stages always have something real to work with.
	if isZero(ctx.Background) {
		ctx.Background = fromHLV(250, 10, 5)
	}
	if isZero(ctx.Foreground) {
		ctx.Foreground = fromHLV(250, 90, 5)
	}
	if isZero(ctx.Accent) {
		// Complement of background hue — at least visually distinct.
		oppositeHue := ctx.Background.Hue() + 180
		if oppositeHue >= 360 {
			oppositeHue -= 360
		}
		ctx.Accent = fromHLV(oppositeHue, 60, 70)
	}

	// Cursor defaults to accent — NeutralizeStage may adjust it.
	ctx.Cursor = ctx.Accent
	return nil
}

// firstTone returns the first non-zero tone found under the given keys.
func firstTone(ctx *Context, keys ...string) core.Tone {
	for _, k := range keys {
		if t, ok := ctx.Palette.Get(k); ok && !isZero(t) {
			return t
		}
	}
	return core.Tone{}
}
