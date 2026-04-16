package sheme

import (
	"github.com/leraniode/wondertone/palette/builtin"
	"testing"
)

func TestGenerate(t *testing.T) {
	p := builtin.Aurora()
	content, err := Generate(p)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	helper := &TestHelper{}
	if !helper.IsValidTheme(content) {
		t.Errorf("Generated theme failed validation:\n%s", content)
	}
}
