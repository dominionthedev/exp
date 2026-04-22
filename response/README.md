# Response (Go)

A Go library for transforming raw technical tool outputs (errors, logs, status) into structured, intelligence-driven responses with human-like qualities. It leverages [illygen](https://github.com/leraniode/illygen) for reasoning and tone selection.

## Core Features

- **Intelligence-Driven**: Uses `illygen` reasoning flows to interpret and rephrase technical data.
- **Structured Outputs**: Each response carries content, tone, confidence scoring, and metadata.
- **Tone Specialization**: Built-in modules for `witty`, `pro`, `casual`, and `joyful` personas.
- **Formatting Adapters**: Built-in support for JSON, Markdown, and ANSI Terminal output.
- **Modular Architecture**: Clean separation between interpretation, tone generation, and formatting.

## Layout

- `core.go`: Foundational types (`Response`, `Emitter`, `Format`).
- `adapter.go`: Bridge between `illygen` results and `response` structures.
- `handler.go`: The `InterpreterNode` which classifies raw technical inputs.
- `format.go`: Formatting logic for different output targets.
- `witty/`, `pro/`, `casual/`, `joyful/`: Specialized reasoning nodes for each tone.

## Usage

### Basic Flow

```go
import (
    "fmt"
    "github.com/dominionthedev/exp/response"
    "github.com/dominionthedev/exp/response/pro"
    illygen "github.com/leraniode/illygen"
)

func main() {
    // 1. Define your reasoning flow
    flow := illygen.NewFlow().
        Add(response.InterpreterNode()).
        Add(pro.NewNode()).
        Link("interpreter", "pro", 1.0)

    engine := illygen.NewEngine()
    adapter := response.NewAdapter(flow, engine)

    // 2. Generate a response from raw input
    resp, err := adapter.Generate(map[string]any{
        "input": "error: database connection timeout",
        "mode":  "pro",
    })

    if err == nil {
        fmt.Println(resp.Content) 
        // Output: "Operation failed with the following error: database connection timeout"
    }
}
```

## Next Steps

- [ ] Implement `illygen` sub-flows for complex reasoning.
- [ ] Add support for template-based response generation in tone nodes.
- [ ] Expand the KnowledgeStore support for domain-specific technical terms.
