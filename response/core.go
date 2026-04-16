package response

import (
	"time"
)

// Response represents a structured, contextual output.
type Response struct {
	// Content is the actual message string.
	Content string `json:"content"`

	// Tone describes the emotional or professional quality of the response.
	Tone string `json:"tone"`

	// Certainty is a metric from 0.0 to 1.0 representing confidence.
	Certainty float64 `json:"certainty"`

	// Metadata contains additional contextual information.
	Metadata map[string]any `json:"metadata,omitempty"`

	// Timestamp is when the response was generated.
	Timestamp time.Time `json:"timestamp"`
}

// New creates a new Response with default values.
func New(content string, tone string) *Response {
	return &Response{
		Content:   content,
		Tone:      tone,
		Certainty: 1.0,
		Metadata:  make(map[string]any),
		Timestamp: time.Now(),
	}
}

// WithCertainty sets the certainty of the response.
func (r *Response) WithCertainty(c float64) *Response {
	r.Certainty = c
	return r
}

// WithMetadata adds a metadata key-value pair to the response.
func (r *Response) WithMetadata(key string, value any) *Response {
	if r.Metadata == nil {
		r.Metadata = make(map[string]any)
	}
	r.Metadata[key] = value
	return r
}

// Emitter is an interface for emitting responses to different sinks.
type Emitter interface {
	Emit(r *Response) error
}

// Format represents the output format for a response.
type Format string

const (
	FormatJSON     Format = "json"
	FormatMarkdown Format = "markdown"
	FormatTerminal Format = "terminal"
)
