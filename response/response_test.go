package response_test

import (
	"testing"

	"github.com/dominionthedev/exp/response"
	"github.com/dominionthedev/exp/response/casual"
	"github.com/dominionthedev/exp/response/pro"
	"github.com/dominionthedev/exp/response/witty"
	illygen "github.com/leraniode/illygen"
)

func TestResponseFlow(t *testing.T) {
	// 1. Setup KnowledgeStore
	store := illygen.NewKnowledgeStore()
	store.Add("witty-success", "witty", map[string]any{
		"category": "success",
		"response": "Boom! Nailed it like a pro.",
	})

	// 2. Setup illygen flow
	flow := illygen.NewFlow().
		Add(response.InterpreterNode()).
		Add(pro.NewNode()).
		Add(witty.NewNode()).
		Add(casual.NewNode()).
		Link("interpreter", "pro", 1.0).
		Link("interpreter", "witty", 1.0).
		Link("interpreter", "casual", 1.0)

	engine := illygen.NewEngine(store)
	adapter := response.NewAdapter(flow, engine)

	// 2. Test professional error response
	t.Run("Professional Error", func(t *testing.T) {
		ctx := map[string]any{
			"input": "error: connection refused",
			"mode":  "pro",
		}
		resp, err := adapter.Generate(ctx)
		if err != nil {
			t.Fatalf("Generate failed: %v", err)
		}

		if resp.Tone != "pro" {
			t.Errorf("expected tone pro, got %s", resp.Tone)
		}
		if resp.Content != "Operation failed with the following error: error: connection refused" {
			t.Errorf("unexpected content: %s", resp.Content)
		}
	})

	// 3. Test witty success response
	t.Run("Witty Success", func(t *testing.T) {
		ctx := map[string]any{
			"input": "success: build finished",
			"mode":  "witty",
		}
		resp, err := adapter.Generate(ctx)
		if err != nil {
			t.Fatalf("Generate failed: %v", err)
		}

		if resp.Tone != "witty" {
			t.Errorf("expected tone witty, got %s", resp.Tone)
		}
		if resp.Content != "Boom! Nailed it like a pro." {
			t.Errorf("unexpected content: %s", resp.Content)
		}
	})

	// 4. Test casual fallback
	t.Run("Casual Info", func(t *testing.T) {
		ctx := map[string]any{
			"input": "server is running",
			"mode":  "casual",
		}
		resp, err := adapter.Generate(ctx)
		if err != nil {
			t.Fatalf("Generate failed: %v", err)
		}

		if resp.Tone != "casual" {
			t.Errorf("expected tone casual, got %s", resp.Tone)
		}
		if resp.Content != "Just so you know: server is running" {
			t.Errorf("unexpected content: %s", resp.Content)
		}
	})
}
