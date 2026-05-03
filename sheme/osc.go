package sheme

import "fmt"

// OSCHandler generates terminal OSC (Operating System Command) escape sequences.
//
// All apply sequences use the format: \033]<code>;<value>\007
// All reset sequences use the format: \033]<reset-code>\007
//
// The hex argument in apply methods accepts:
//   - A shell variable reference: "$COLOR_FG"
//   - A literal hex colour:       "#1a1a2e"
//
// Compatibility: xterm, iTerm2, WezTerm, Alacritty, Kitty, foot, tmux (passthrough).
// Not supported by: Windows Console, very old VTE versions.
type OSCHandler struct{}

// ── Apply sequences ───────────────────────────────────────────────────────────

// Foreground emits OSC 10 — set default text foreground colour.
func (h *OSCHandler) Foreground(hex string) string {
	return fmt.Sprintf("\\033]10;%s\\007", hex)
}

// Background emits OSC 11 — set default background colour.
func (h *OSCHandler) Background(hex string) string {
	return fmt.Sprintf("\\033]11;%s\\007", hex)
}

// Cursor emits OSC 12 — set cursor colour.
func (h *OSCHandler) Cursor(hex string) string {
	return fmt.Sprintf("\\033]12;%s\\007", hex)
}

// ANSI emits OSC 4 — set a colour in the ANSI 0–255 palette.
func (h *OSCHandler) ANSI(index int, hex string) string {
	return fmt.Sprintf("\\033]4;%d;%s\\007", index, hex)
}

// ── Reset sequences ───────────────────────────────────────────────────────────
// Used by the reset tool to restore the terminal to its default colours
// without needing to know what the previous colours were.

// ResetANSI emits OSC 104 — reset a single ANSI colour index to terminal default.
// Pass -1 to reset all ANSI colours at once.
func (h *OSCHandler) ResetANSI(index int) string {
	if index < 0 {
		return "\\033]104\\007" // reset all
	}
	return fmt.Sprintf("\\033]104;%d\\007", index)
}

// ResetForeground emits OSC 110 — restore foreground to terminal default.
func (h *OSCHandler) ResetForeground() string { return "\\033]110\\007" }

// ResetBackground emits OSC 111 — restore background to terminal default.
func (h *OSCHandler) ResetBackground() string { return "\\033]111\\007" }

// ResetCursor emits OSC 112 — restore cursor colour to terminal default.
func (h *OSCHandler) ResetCursor() string { return "\\033]112\\007" }

// ResetAll returns the complete set of sequences needed to undo a theme.
// Intended for use by reset tools; returns one sequence per line.
func (h *OSCHandler) ResetAll() string {
	return h.ResetForeground() + "\n" +
		h.ResetBackground() + "\n" +
		h.ResetCursor() + "\n" +
		h.ResetANSI(-1) // resets all 256 ANSI slots in one shot
}
