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

func TestToTitle(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"the quick brown fox", "The Quick Brown Fox"},
		{"ALREADY UPPER", "Already Upper"},
	}
	for _, tt := range tests {
		got := ToTitle(tt.input)
		if got != tt.expected {
			t.Errorf("ToTitle(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}

func TestToPascal_Unicode(t *testing.T) {
	// Ensure non-ASCII letters are handled without panic
	got := ToPascal("café latte")
	if got == "" {
		t.Error("ToPascal returned empty for unicode input")
	}
}

func TestNewUtilities(t *testing.T) {
	if Lower("HELLO") != "hello" {
		t.Error("Lower failed")
	}
	if Upper("hello") != "HELLO" {
		t.Error("Upper failed")
	}
	if SwapCase("Hello World") != "hELLO wORLD" {
		t.Error("SwapCase failed")
	}
	if !StartsWith("hello", "he") {
		t.Error("StartsWith failed")
	}
	if !EndsWith("hello", "lo") {
		t.Error("EndsWith failed")
	}
	if !Contains("hello", "ell") {
		t.Error("Contains failed")
	}
	if CountOccurrences("hello", "l") != 2 {
		t.Error("CountOccurrences failed")
	}
	if !IsPalindrome("A man, a plan, a canal: Panama") {
		t.Error("IsPalindrome failed")
	}
	if MaskText("1234567890", 2, 2, "*") != "12******90" {
		t.Errorf("MaskText failed: %s", MaskText("1234567890", 2, 2, "*"))
	}
	if FirstWord("  hello world  ") != "hello" {
		t.Error("FirstWord failed")
	}
	if LastWord("hello world  ") != "world" {
		t.Error("LastWord failed")
	}
	if len(RandomString(10, "")) != 10 {
		t.Error("RandomString failed")
	}
}
