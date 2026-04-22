package pro

import (
	illygen "github.com/leraniode/illygen"
)

// NewNode creates a professional response node that transforms technical inputs.
func NewNode() *illygen.Node {
	return illygen.NewNode("pro", func(ctx illygen.Context) illygen.Result {
		ks := illygen.Knowledge(ctx)
		category := ctx.String("category")
		input := ctx.String("input")

		ctx.Set("tone", "pro")

		if ks != nil {
			for _, u := range ks.Domain("pro") {
				if u.Facts["category"] == category {
					return illygen.Result{
						Value:      u.Facts["response"],
						Confidence: u.Weight,
					}
				}
			}
		}

		// Fallback for professional formal tone
		switch category {
		case "error":
			return illygen.Result{Value: "Operation failed with the following error: " + input, Confidence: 1.0}
		case "success":
			return illygen.Result{Value: "Action successfully completed: " + input, Confidence: 1.0}
		case "warning":
			return illygen.Result{Value: "Caution: " + input, Confidence: 0.8}
		}

		return illygen.Result{Value: "Reported: " + input, Confidence: 1.0}
	})
}
