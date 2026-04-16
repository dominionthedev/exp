package strings

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		input     string
		separator string
		expected  string
	}{
		{"Hello World!", "-", "hello-world"},
		{"Some_Cool_Text", "_", "some_cool_text"},
		{"  Trim Me  ", "-", "trim-me"},
	}

	for _, test := range tests {
		result := Slugify(test.input, test.separator)
		if result != test.expected {
			t.Errorf("Slugify(%q, %q) = %q; expected %q", test.input, test.separator, result, test.expected)
		}
	}
}

func TestTruncate(t *testing.T) {
	if Truncate("Hello World", 5, "...") != "Hello..." {
		t.Error("Truncate failed")
	}
}

func TestCases(t *testing.T) {
	text := "hello world_this-is PascalCase"
	if ToSnake(text) != "hello_world_this_is_pascal_case" {
		t.Errorf("ToSnake failed: %s", ToSnake(text))
	}
	if ToKebab(text) != "hello-world-this-is-pascal-case" {
		t.Errorf("ToKebab failed: %s", ToKebab(text))
	}
	if ToPascal("hello world") != "HelloWorld" {
		t.Errorf("ToPascal failed: %s", ToPascal("hello world"))
	}
	if ToCamel("hello world") != "helloWorld" {
		t.Errorf("ToCamel failed: %s", ToCamel("hello world"))
	}
}

func TestUtilities(t *testing.T) {
	if Capitalize("hello") != "Hello" {
		t.Error("Capitalize failed")
	}
	if Reverse("abc") != "cba" {
		t.Error("Reverse failed")
	}
	if Repeat("a", 3) != "aaa" {
		t.Error("Repeat failed")
	}
	if PadLeft("a", 3, " ") != "  a" {
		t.Errorf("PadLeft failed: %q", PadLeft("a", 3, " "))
	}
	if PadRight("a", 3, " ") != "a  " {
		t.Errorf("PadRight failed: %q", PadRight("a", 3, " "))
	}
	if Pad("a", 3, " ") != " a " {
		t.Errorf("Pad failed: %q", Pad("a", 3, " "))
	}
	if StripPrefix("prefix_hello", "prefix_") != "hello" {
		t.Error("StripPrefix failed")
	}
	if StripSuffix("hello_suffix", "_suffix") != "hello" {
		t.Error("StripSuffix failed")
	}
	if WordCount("hello world again") != 3 {
		t.Error("WordCount failed")
	}
	if SentenceCount("Hello! How are you? Fine.") != 3 {
		t.Error("SentenceCount failed")
	}
}
