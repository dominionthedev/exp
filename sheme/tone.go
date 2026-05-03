package sheme

import (
	"math"

	"github.com/leraniode/wondertone/core"
)

// ── Tone construction helpers ─────────────────────────────────────────────────

// withL returns a new tone with lightness replaced by l, H and V preserved.
func withL(t core.Tone, l float64) core.Tone {
	return core.New(core.Light(clamp(l, 0, 100)), core.Vibrancy(t.Vibrancy()), core.Hue(t.Hue()))
}

// withV returns a new tone with vibrancy replaced by v, H and L preserved.
func withV(t core.Tone, v float64) core.Tone {
	return core.New(core.Light(t.Light()), core.Vibrancy(clamp(v, 0, 100)), core.Hue(t.Hue()))
}

// withLV returns a new tone with both L and V replaced, H preserved.
func withLV(t core.Tone, l, v float64) core.Tone {
	return core.New(core.Light(clamp(l, 0, 100)), core.Vibrancy(clamp(v, 0, 100)), core.Hue(t.Hue()))
}

// fromHLV builds a tone from explicit hue, lightness, and vibrancy values.
func fromHLV(h, l, v float64) core.Tone {
	return core.New(core.Light(clamp(l, 0, 100)), core.Vibrancy(clamp(v, 0, 100)), core.Hue(h))
}

// clampTone clamps L and V into [minL, maxL] and [minV, maxV], hue preserved.
func clampTone(t core.Tone, minL, maxL, minV, maxV float64) core.Tone {
	return core.New(
		core.Light(clamp(t.Light(), minL, maxL)),
		core.Vibrancy(clamp(t.Vibrancy(), minV, maxV)),
		core.Hue(t.Hue()),
	)
}

// isZero reports whether t is the zero value (unset tone).
func isZero(t core.Tone) bool { return t == (core.Tone{}) }

// ── Measurement helpers ───────────────────────────────────────────────────────

// deltaL returns the absolute lightness difference between two tones.
func deltaL(a, b core.Tone) float64 { return math.Abs(a.Light() - b.Light()) }

// deltaH returns the minimum angular hue distance between two tones (0–180°).
func deltaH(a, b core.Tone) float64 {
	d := math.Abs(a.Hue() - b.Hue())
	if d > 180 {
		d = 360 - d
	}
	return d
}

// relativeLuminance returns an sRGB approximation of perceptual luminance
// for a Wondertone light value (0–100 scale → 0.0–1.0).
func relativeLuminance(lightness float64) float64 {
	l := lightness / 100.0
	if l <= 0.03928 {
		return l / 12.92
	}
	return math.Pow((l+0.055)/1.055, 2.4)
}

// contrastRatio returns the WCAG contrast ratio between two lightness values.
// A ratio ≥ 4.5 satisfies WCAG AA for normal text.
func contrastRatio(l1, l2 float64) float64 {
	lum1 := relativeLuminance(l1)
	lum2 := relativeLuminance(l2)
	if lum1 < lum2 {
		lum1, lum2 = lum2, lum1
	}
	return (lum1 + 0.05) / (lum2 + 0.05)
}

// ── General math ─────────────────────────────────────────────────────────────

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func lerp(a, b, t float64) float64 { return a + (b-a)*t }
