# Sheme

A Go library for creating shell themes from [Wondertone](https://github.com/leraniode/wondertone) palettes. It generates bash/zsh compatible `.theme` files containing both OSC sequences and shell-syntax color tokens.

## Overview

Sheme (pronounced "scheme") bridge the gap between perceptual color palettes and your terminal environment. A **tone** is a single color, while a **palette** is a complete theme (`.theme` file) with multiple tones.

## Features

- **Wondertone Integration**: Full support for `.wtone` files and built-in palettes (Aurora, Midnight, etc.).
- **Dual Format**: Generates `.theme` files that include both `export` variables for shell scripts and `printf` OSC sequences for immediate terminal application.
- **Modular API**: Clean separation between OSC generation, shell compatibility, and theme management.
- **Validation**: Built-in helpers to verify the integrity of generated themes.

## Installation

```bash
go get github.com/dominionthedev/exp/sheme
```

## Usage

### As a Library

```go
import (
    "fmt"
    "github.com/dominionthedev/exp/sheme"
    "github.com/leraniode/wondertone/palette/builtin"
)

func main() {
    p := builtin.Aurora()
    content, err := sheme.Generate(p)
    if err != nil {
        panic(err)
    }
    fmt.Print(content)
}
```

### The .theme Format

The generated `.theme` file looks like this:

```bash
# Theme: Aurora
# Author: leraniode

# Color Tokens
export COLOR_FG="#0d1213"
export COLOR_BG="#f1f4f4"
...

# OSC Sequences
printf "\033]10;$COLOR_FG\007"
printf "\033]11;$COLOR_BG\007"
...
```

## Modules

- `generator.go`: Logic for mapping Wondertone palettes to terminal ANSI indexes.
- `osc.go`: Low-level OSC sequence generation.
- `shell.go`: Shell-specific syntax and variable export logic.
- `theme.go`: Data structures for themes and serialization.
- `sheme.go`: High-level API for easy theme generation.
- `test.go`: Validation and testing utilities.

## License

MIT
