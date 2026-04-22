package joyful

import (
	illygen "github.com/leraniode/illygen"
)

// NewNode creates a joyful response node that transforms technical inputs.
func NewNode() *illygen.Node {
	return illygen.NewNode("joyful", func(ctx illygen.Context) illygen.Result {
		ks := illygen.Knowledge(ctx)
		category := ctx.String("category")
		input := ctx.String("input")

		ctx.Set("tone", "joyful")

		if ks != nil {
			for _, u := range ks.Domain("joyful") {
				if u.Facts["category"] == category {
					return illygen.Result{
						Value:      u.Facts["response"],
						Confidence: u.Weight,
					}
				}
			}
		}

		// Fallback for joyful enthusiastic tone
		switch category {
		case "error":
			return illygen.Result{Value: "We encountered a little bump, but we'll fix it together! Error: " + input, Confidence: 0.7}
		case "success":
			return illygen.Result{Value: "Hooray! Everything worked perfectly! " + input, Confidence: 1.0}
		case "warning":
			return illygen.Result{Value: "Just a small heads-up, but nothing to worry about! " + input, Confidence: 0.9}
		}

		return illygen.Result{Value: "I'm so happy to share this with you: " + input, Confidence: 1.0}
	})
}
