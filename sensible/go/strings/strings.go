package strings

import (
	"regexp"
	"strings"
)

var (
	reNonAlphaNum = regexp.MustCompile(`[^\w\s-]`)
	reSpaces      = regexp.MustCompile(`[-\s]+`)
	reUpperLower  = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)

// Slugify converts text to a URL-friendly slug.
// Simplified version for standard library.
func Slugify(text string, separator string) string {
	result := strings.ToLower(text)
	result = reNonAlphaNum.ReplaceAllString(result, "")
	result = reSpaces.ReplaceAllString(result, separator)
	return strings.Trim(result, separator)
}

// Truncate truncates text to a specific length with an optional suffix.
func Truncate(text string, length int, suffix string) string {
	if len(text) <= length {
		return text
	}
	return strings.TrimSpace(text[:length]) + suffix
}

// ToSnake converts text to snake_case.
func ToSnake(text string) string {
	s := reUpperLower.ReplaceAllString(text, "${1}_${2}")
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[-_\s]+`).ReplaceAllString(s, "_")
	return strings.Trim(s, "_")
}

// ToPascal converts text to PascalCase.
func ToPascal(text string) string {
	words := regexp.MustCompile(`[-_\s]+`).Split(text, -1)
	var result string
	for _, word := range words {
		if word != "" {
			// Title is deprecated but good for simple cases. 
			// For a production lib we'd use cases package or manual capitalization.
			result += strings.Title(strings.ToLower(word))
		}
	}
	return result
}

// ToCamel converts text to camelCase.
func ToCamel(text string) string {
	s := ToPascal(text)
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// ToKebab converts text to kebab-case.
func ToKebab(text string) string {
	return strings.ReplaceAll(ToSnake(text), "_", "-")
}

// Capitalize capitalizes the first letter of the text.
func Capitalize(text string) string {
	if text == "" {
		return ""
	}
	return strings.ToUpper(text[:1]) + text[1:]
}

// Reverse reverses the text.
func Reverse(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Repeat repeats the text a given number of times.
func Repeat(text string, times int) string {
	return strings.Repeat(text, times)
}

// PadLeft pads text on the left to reach a certain length.
func PadLeft(text string, length int, char string) string {
	if len(text) >= length {
		return text
	}
	return strings.Repeat(char, length-len(text)) + text
}

// PadRight pads text on the right to reach a certain length.
func PadRight(text string, length int, char string) string {
	if len(text) >= length {
		return text
	}
	return text + strings.Repeat(char, length-len(text))
}

// Pad pads text on both sides to reach a certain length.
func Pad(text string, length int, char string) string {
	if len(text) >= length {
		return text
	}
	padding := length - len(text)
	left := padding / 2
	right := padding - left
	return strings.Repeat(char, left) + text + strings.Repeat(char, right)
}

// StripPrefix removes a prefix from the text if it exists.
func StripPrefix(text string, prefix string) string {
	return strings.TrimPrefix(text, prefix)
}

// StripSuffix removes a suffix from the text if it exists.
func StripSuffix(text string, suffix string) string {
	return strings.TrimSuffix(text, suffix)
}

// WordCount counts the number of words in text.
func WordCount(text string) int {
	return len(strings.Fields(text))
}

// SentenceCount counts the number of sentences in text.
func SentenceCount(text string) int {
	return len(regexp.MustCompile(`[^.!?]+[.!?]+`).FindAllString(text, -1))
}
