package witty

import (
	illygen "github.com/leraniode/illygen"
)

func NewNode() *illygen.Node {
	return illygen.NewNode("witty", func(ctx illygen.Context) illygen.Result {
		ks := illygen.Knowledge(ctx)
		category := ctx.String("category")
		
		if ks != nil {
			for _, u := range ks.Domain("witty") {
				if u.Facts["category"] == category {
					ctx.Set("tone", "witty")
					return illygen.Result{Value: u.Facts["response"], Confidence: u.Weight}
				}
			}
		}
		return illygen.Result{Value: "Something weird happened: " + ctx.String("input"), Confidence: 0.5}
	})
}
