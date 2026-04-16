package sheme

import "fmt"

// OSCHandler handles terminal OSC sequence generation.
type OSCHandler struct{}

// Foreground returns the OSC sequence for setting the foreground color.
func (h *OSCHandler) Foreground(hex string) string {
	return fmt.Sprintf("\\033]10;%s\\007", hex)
}

// Background returns the OSC sequence for setting the background color.
func (h *OSCHandler) Background(hex string) string {
	return fmt.Sprintf("\\033]11;%s\\007", hex)
}

// Cursor returns the OSC sequence for setting the cursor color.
func (h *OSCHandler) Cursor(hex string) string {
	return fmt.Sprintf("\\033]12;%s\\007", hex)
}

// Color returns the OSC sequence for setting an ANSI color index.
func (h *OSCHandler) Color(index int, hex string) string {
	return fmt.Sprintf("\\033]4;%d;%s\\007", index, hex)
}
