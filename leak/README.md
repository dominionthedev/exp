# Leak

A Python module for terminal service management and event-driven terminal interactions.

## Features

- **Event Bus**: Asynchronous, `asyncio`-based event dispatcher.
- **Service Management**: Lifecycle control (`start`/`stop`) and OS signal handling (SIGINT, SIGTERM, SIGWINCH).
- **Terminal Utilities**: Helpers for raw mode, cursor management, and terminal sizing.

## Installation

```bash
uv add leak
```

## Usage

### Simple Service

```python
import asyncio
from leak.service import Service
from leak.events import Event

async def main():
    service = Service()
    
    async def on_start(event):
        print("Service started!")
        
    service.events.subscribe("service_start", on_start)
    
    await service.run_forever()

if __name__ == "__main__":
    asyncio.run(main())
```

### Raw Mode

```python
from leak.terminal import Terminal

with Terminal.raw_mode():
    print("In raw mode!")
    # Read raw input...
```

## License

MIT
