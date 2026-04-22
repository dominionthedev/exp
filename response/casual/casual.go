package casual

import (
	illygen "github.com/leraniode/illygen"
)

// NewNode creates a casual response node that transforms technical inputs.
func NewNode() *illygen.Node {
	return illygen.NewNode("casual", func(ctx illygen.Context) illygen.Result {
		ks := illygen.Knowledge(ctx)
		category := ctx.String("category")
		input := ctx.String("input")

		ctx.Set("tone", "casual")

		if ks != nil {
			for _, u := range ks.Domain("casual") {
				if u.Facts["category"] == category {
					return illygen.Result{
						Value:      u.Facts["response"],
						Confidence: u.Weight,
					}
				}
			}
		}

		// Fallback for casual conversational tone
		switch category {
		case "error":
			return illygen.Result{Value: "Wait, that didn't work. Problem was: " + input, Confidence: 0.8}
		case "success":
			return illygen.Result{Value: "Done! No sweat.", Confidence: 1.0}
		}

		return illygen.Result{Value: "Just so you know: " + input, Confidence: 1.0}
	})
}
