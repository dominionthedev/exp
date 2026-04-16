package sheme

import (
	"fmt"
)

// ShellHandler handles shell-specific syntax and exports.
type ShellHandler struct{}

// Variable returns a shell variable assignment.
func (h *ShellHandler) Variable(name, value string) string {
	return fmt.Sprintf("export %s=\"%s\"", name, value)
}

// Printf returns a shell printf command for an OSC sequence.
func (h *ShellHandler) Printf(sequence string) string {
	return fmt.Sprintf("printf \"%s\"", sequence)
}
