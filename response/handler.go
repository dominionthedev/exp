package response

import (
	illygen "github.com/leraniode/illygen"
	"strings"
)

// InterpreterNode analyzes raw input (logs, errors) and classifies it.
func InterpreterNode() *illygen.Node {
	return illygen.NewNode("interpreter", func(ctx illygen.Context) illygen.Result {
		input := strings.ToLower(ctx.String("input"))
		
		// Basic classification logic
		category := "info"
		if strings.Contains(input, "error") || strings.Contains(input, "fail") || strings.Contains(input, "panic") {
			category = "error"
		} else if strings.Contains(input, "warn") || strings.Contains(input, "limit") {
			category = "warning"
		} else if strings.Contains(input, "success") || strings.Contains(input, "done") || strings.Contains(input, "ok") {
			category = "success"
		}

		ctx.Set("category", category)
		
		// Route to the requested tone (defaulting to pro if not specified)
		mode := ctx.String("mode")
		if mode == "" {
			mode = "pro"
		}
		
		return illygen.Result{Next: mode, Confidence: 1.0}
	})
}
