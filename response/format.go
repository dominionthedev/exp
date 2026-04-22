package response

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ToJSON converts a response to a JSON string.
func ToJSON(r *Response) (string, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ToMarkdown converts a response to a Markdown formatted string.
func ToMarkdown(r *Response) string {
	return fmt.Sprintf("### [%s] Response\n\n**Content:** %s\n\n**Confidence:** %.2f\n\n*Generated at: %s*", 
		r.Tone, r.Content, r.Certainty, r.Timestamp.Format("2006-01-02 15:04:05"))
}

// ToTerminal converts a response to an ANSI-colored terminal string.
func ToTerminal(r *Response) string {
	color := "\033[0m" // Reset
	switch r.Tone {
	case "witty":
		color = "\033[35m" // Magenta
	case "pro":
		color = "\033[34m" // Blue
	case "casual":
		color = "\033[32m" // Green
	case "error":
		color = "\033[31m" // Red
	}
	
	return fmt.Sprintf("%s[%s] %s\033[0m", color, strings.ToUpper(r.Tone), r.Content)
}
