# Response (Go)

A Go library for structured, intelligence-driven tool outputs from plain error messages, infos and logs. Instead of returning plain strings, **response** provides rich, contextual outputs with emotions, tones and human-like qualities.

It integrates with [illygen](https://github.com/leraniode/illygen) for building reasoning-based response flows.

## Core Features

- **Structured Outputs**: Define complex response types beyond simple strings
- **Confidence Scoring**: Each response can carry a "certainty" metric (0.0-1.0)
- **Metadata Support**: Attach contextual data to any response
- **Illygen Integration**: Use illygen flows to generate the "best" response
- **Formatting Adapters**: Export responses to JSON, Markdown, or Terminal (OSC/ANSI)

## Types

There are different types of responses `response` provides:

- [witty](./witty/): For humorous or clever outputs
- [pro](./pro/): For professional, formal outputs
- [casual](./casual/): For informal, conversational outputs
- [joyful](./joyful/): For happy, enthusiastic outputs

## Layout

- core: the core types and functions for building responses
- adapter: the module for using illygen to generate responses
- handler: the module for handling the input meesages
- format: the module for formatting responses
- witty/: the module for witty responses
- pro/: the module for professional responses
- casual/: the module for casual responses
- joyful/: the module for joyful responses

## Building

You can build your own response types using the core types and functions. For example:

