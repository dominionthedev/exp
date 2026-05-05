package strings

import (
	"math/rand"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	reNonAlphaNum = regexp.MustCompile(`[^\w\s-]`)
	reSpaces      = regexp.MustCompile(`[-\s]+`)
	reUpperLower  = regexp.MustCompile(`([a-z0-9])([A-Z])`)
	reSplit       = regexp.MustCompile(`[-_\s]+`)

	// titler is the shared Unicode-correct title caser (language-neutral).
	titler = cases.Title(language.Und, cases.NoLower)
)

// capitaliseWord upper-cases the first rune of a word, lower-cases the rest.
func capitaliseWord(w string) string {
	runes := []rune(w)
	if len(runes) == 0 {
		return ""
	}
	return string(unicode.ToUpper(runes[0])) + strings.ToLower(string(runes[1:]))
}

// Slugify converts text to a URL-friendly slug.
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
	s = reSplit.ReplaceAllString(s, "_")
	return strings.Trim(s, "_")
}

// ToPascal converts text to PascalCase using Unicode-correct capitalisation.
func ToPascal(text string) string {
	words := reSplit.Split(text, -1)
	var result string
	for _, word := range words {
		if word != "" {
			result += capitaliseWord(word)
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
	runes := []rune(s)
	return string(unicode.ToLower(runes[0])) + string(runes[1:])
}

// ToKebab converts text to kebab-case.
func ToKebab(text string) string {
	return strings.ReplaceAll(ToSnake(text), "_", "-")
}

// ToTitle converts text to Title Case using Unicode-correct rules.
func ToTitle(text string) string {
	return titler.String(strings.ToLower(text))
}

// Capitalize upper-cases the first Unicode letter of text, preserving the rest.
func Capitalize(text string) string {
	runes := []rune(text)
	if len(runes) == 0 {
		return ""
	}
	return string(unicode.ToUpper(runes[0])) + string(runes[1:])
}

// Reverse reverses text, correctly handling multi-byte Unicode characters.
func Reverse(text string) string {
	runes := []rune(text)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Repeat repeats text a given number of times.
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

// StripPrefix removes a prefix from text if it exists.
func StripPrefix(text string, prefix string) string {
	return strings.TrimPrefix(text, prefix)
}

// StripSuffix removes a suffix from text if it exists.
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

// Lower converts text to lowercase.
func Lower(text string) string {
	return strings.ToLower(text)
}

// Upper converts text to uppercase.
func Upper(text string) string {
	return strings.ToUpper(text)
}

// SwapCase swaps the case of all letters in the text.
func SwapCase(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, text)
}

// StartsWith checks if text starts with a prefix.
func StartsWith(text string, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}

// EndsWith checks if text ends with a suffix.
func EndsWith(text string, suffix string) bool {
	return strings.HasSuffix(text, suffix)
}

// Contains checks if text contains a substring.
func Contains(text string, substring string) bool {
	return strings.Contains(text, substring)
}

// CountOccurrences counts occurrences of a substring in text.
func CountOccurrences(text string, substring string) int {
	return strings.Count(text, substring)
}

// RandomString generates a random string of a given length.
func RandomString(length int, charset string) string {
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// IsPalindrome checks if text is a palindrome (ignoring case and non-alphanumeric characters).
func IsPalindrome(text string) bool {
	var cleaned []rune
	for _, r := range strings.ToLower(text) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
			return false
		}
	}
	return true
}

// MaskText masks text, showing only the specified number of characters at start and end.
func MaskText(text string, visibleStart int, visibleEnd int, maskChar string) string {
	runes := []rune(text)
	if len(runes) <= visibleStart+visibleEnd {
		return text
	}
	maskedLen := len(runes) - visibleStart - visibleEnd
	result := string(runes[:visibleStart])
	result += strings.Repeat(maskChar, maskedLen)
	if visibleEnd > 0 {
		result += string(runes[len(runes)-visibleEnd:])
	}
	return result
}

// FirstWord gets the first word from text.
func FirstWord(text string) string {
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return ""
	}
	return fields[0]
}

// LastWord gets the last word from text.
func LastWord(text string) string {
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return ""
	}
	return fields[len(fields)-1]
}
