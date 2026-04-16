package chains

import (
	"strings"
	"testing"
)

func TestChain(t *testing.T) {
	t.Run("Map", func(t *testing.T) {
		res := New(10).
			Map(func(x int) int { return x + 5 }).
			Map(func(x int) int { return x * 2 }).
			Value()

		if res != 30 {
			t.Errorf("expected 30, got %d", res)
		}
	})

	t.Run("Then", func(t *testing.T) {
		c1 := New("  hello  ")
		c2 := Then(c1, strings.TrimSpace)
		c3 := Then(c2, strings.ToUpper)
		res := c3.Value()

		if res != "HELLO" {
			t.Errorf("expected HELLO, got %s", res)
		}
	})

	t.Run("Tap", func(t *testing.T) {
		var sideEffect int
		res := New(10).
			Tap(func(x int) { sideEffect = x }).
			Map(func(x int) int { return x * 2 }).
			Value()

		if res != 20 {
			t.Errorf("expected 20, got %d", res)
		}
		if sideEffect != 10 {
			t.Errorf("expected sideEffect 10, got %d", sideEffect)
		}
	})

	t.Run("Pipe", func(t *testing.T) {
		add5 := func(x int) int { return x + 5 }
		mul2 := func(x int) int { return x * 2 }
		res := New(10).Pipe(add5, mul2).Value()

		if res != 30 {
			t.Errorf("expected 30, got %d", res)
		}
	})
}
