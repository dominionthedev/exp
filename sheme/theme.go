package sheme

import (
	"fmt"
	"github.com/leraniode/wondertone/core"
	"strings"
)

// Theme represents a standard terminal theme.
type Theme struct {
	Name       string
	Author     string
	Foreground core.Tone
	Background core.Tone
	Cursor     core.Tone

	// ANSI 0-15
	Colors [16]core.Tone
}

// Export returns the complete .theme file content including variables and OSC sequences.
func (t *Theme) Export() string {
	var sb strings.Builder
	osc := &OSCHandler{}
	sh := &ShellHandler{}

	sb.WriteString(fmt.Sprintf("# Theme: %s\n", t.Name))
	sb.WriteString(fmt.Sprintf("# Author: %s\n\n", t.Author))

	sb.WriteString("# Color Tokens\n")
	sb.WriteString(sh.Variable("COLOR_FG", t.Foreground.Hex()) + "\n")
	sb.WriteString(sh.Variable("COLOR_BG", t.Background.Hex()) + "\n")
	sb.WriteString(sh.Variable("COLOR_CURSOR", t.Cursor.Hex()) + "\n")

	for i, tone := range t.Colors {
		sb.WriteString(sh.Variable(fmt.Sprintf("COLOR_%d", i), tone.Hex()) + "\n")
	}
	sb.WriteString("\n")

	sb.WriteString("# OSC Sequences\n")
	sb.WriteString(sh.Printf(osc.Foreground("$COLOR_FG")) + "\n")
	sb.WriteString(sh.Printf(osc.Background("$COLOR_BG")) + "\n")
	sb.WriteString(sh.Printf(osc.Cursor("$COLOR_CURSOR")) + "\n")

	for i := range t.Colors {
		sb.WriteString(sh.Printf(osc.Color(i, fmt.Sprintf("$COLOR_%d", i))) + "\n")
	}

	return sb.String()
}
