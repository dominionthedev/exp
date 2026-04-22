package response

import (
	"fmt"
	illygen "github.com/leraniode/illygen"
)

// Adapter connects illygen reasoning flows to the response system.
type Adapter struct {
	flow   *illygen.Flow
	engine *illygen.Engine
}

// NewAdapter creates a new response adapter with a flow and engine.
func NewAdapter(flow *illygen.Flow, engine *illygen.Engine) *Adapter {
	return &Adapter{
		flow:   flow,
		engine: engine,
	}
}

// Generate runs the underlying illygen flow and converts the result to a Response.
func (a *Adapter) Generate(ctx map[string]any) (*Response, error) {
	illyCtx := illygen.Context{}
	for k, v := range ctx {
		illyCtx[k] = v
	}

	res, err := a.engine.Run(a.flow, illyCtx)
	if err != nil {
		return nil, err
	}

	resp := New(fmt.Sprintf("%v", res.Value), "neutral")
	resp.WithCertainty(res.Confidence)

	// If the nodes updated the context with a specific tone, use it.
	if tone := illyCtx.String("tone"); tone != "" {
		resp.Tone = tone
	}

	return resp, nil
}
