package sheme

import (
	"fmt"
	"strings"

	"github.com/leraniode/wondertone/core"
)

// Theme is the fully resolved color set for a terminal theme.
//
// It is pure data — it knows nothing about files, paths, or shell state.
// Writing, naming, or sourcing the output is the caller's responsibility.
type Theme struct {
	Name       string
	Author     string
	Background core.Tone
	Foreground core.Tone
	Cursor     core.Tone
	Colors     [16]core.Tone // ANSI 0–15
}

// Render serializes the Theme to a shell-sourceable string.
//
// The output is POSIX-compatible and suitable for bash/zsh. It contains:
//  1. A header comment with name and author
//  2. COLOR_FG / COLOR_BG / COLOR_CURSOR exports
//  3. COLOR_0 through COLOR_15 exports (ANSI palette)
//  4. OSC sequences to repaint the terminal emulator live
//
// Nothing is written to disk — the caller decides what to do with the string.
func (t *Theme) Render() string {
	var sb strings.Builder
	osc := &OSCHandler{}
	sh := &ShellHandler{}

	// Header
	sb.WriteString(fmt.Sprintf("# Theme: %s\n", t.Name))
	if t.Author != "" {
		sb.WriteString(fmt.Sprintf("# Author: %s\n", t.Author))
	}
	sb.WriteString("\n")

	// Base color tokens
	sb.WriteString("# ── Color tokens ────────────────────────────────────────\n")
	sb.WriteString(sh.Export("COLOR_FG", t.Foreground.Hex()) + "\n")
	sb.WriteString(sh.Export("COLOR_BG", t.Background.Hex()) + "\n")
	sb.WriteString(sh.Export("COLOR_CURSOR", t.Cursor.Hex()) + "\n\n")

	// ANSI palette tokens
	for i, tone := range t.Colors {
		sb.WriteString(sh.Export(fmt.Sprintf("COLOR_%d", i), tone.Hex()) + "\n")
	}
	sb.WriteString("\n")

	// OSC sequences — repaint the terminal emulator live when sourced
	sb.WriteString("# ── Terminal sequences ──────────────────────────────────\n")
	sb.WriteString(sh.Print(osc.Foreground("$COLOR_FG")) + "\n")
	sb.WriteString(sh.Print(osc.Background("$COLOR_BG")) + "\n")
	sb.WriteString(sh.Print(osc.Cursor("$COLOR_CURSOR")) + "\n")
	for i := range t.Colors {
		sb.WriteString(sh.Print(osc.ANSI(i, fmt.Sprintf("$COLOR_%d", i))) + "\n")
	}

	return sb.String()
}
