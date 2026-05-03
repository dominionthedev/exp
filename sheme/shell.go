package sheme

import "fmt"

// ShellHandler produces POSIX-compatible shell syntax fragments.
// It is intentionally minimal — only the two constructs sheme actually needs.
type ShellHandler struct{}

// Export returns a shell variable export statement.
//
//	Export("COLOR_FG", "#1a1a2e") → `export COLOR_FG="#1a1a2e"`
func (h *ShellHandler) Export(name, value string) string {
	return fmt.Sprintf(`export %s="%s"`, name, value)
}

// Print returns a printf statement that emits an OSC escape sequence.
//
//	Print("\\033]10;#fff\\007") → `printf "\\033]10;#fff\\007"`
func (h *ShellHandler) Print(seq string) string {
	return fmt.Sprintf(`printf "%s"`, seq)
}
