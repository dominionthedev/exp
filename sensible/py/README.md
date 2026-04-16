# Sensible (Python)

A Python library of simple and smart logics, types, functions, classes, chains(loops), and decorators for data modeling, workflows and development.

## Features

### Types (Result/Option Patterns)

Safe error handling and optional value management:

```python
from sensible.types import Result, Option

# Result for error handling
res = Result.ok(10).map(lambda x: x * 2).map(lambda x: x + 5)
print(res.unwrap())  # 25

err = Result.error("failed").and_then(lambda x: Result.ok(x * 2))
print(err.error_value())  # "failed"

# Option for optional values
opt = Option.some(42).filter(lambda x: x > 10)
print(opt.unwrap())  # 42

none_val = Option.none().unwrap_or(0)  # Default if None
```

### Chains (Fluent Composition)

Compose operations in a fluent, readable way:

```python
from sensible.chains import chain

result = (
    chain("  hello  ")
    .then(str.strip)
    .then(str.upper)
    .then(lambda s: f"[{s}]")
    .value()
)  # "[HELLO]"

# Chain with conditional logic
result = (
    chain(None)
    .if_none("default")
    .then(str.upper)
    .value()
)  # "DEFAULT"
```

### Decorators

Powerful decorators for common patterns:

```python
from sensible.decorators import cached, lru_cache, retry, validate_args, timed

# Caching with size limits
@cached(max_size=100)
def expensive_func(x):
    return x * 2

# LRU caching
@lru_cache(max_size=50)
def compute(x):
    return x ** 2

# Retry with delay
@retry(retries=3, delay=1.0, exceptions=ValueError)
def risky_op():
    ...

# Type validation
@validate_args(str, int)
def greet(name: str, age: int):
    return f"{name} is {age}"

# Timing execution
@timed
def slow_func():
    ...
```

### String Utilities

Common string operations:

```python
from sensible.strings import (
    slugify, truncate, capitalize, pad, pluralize,
    to_snake, to_camel, to_pascal, mask_text, random_string
)

# Case conversions
to_snake("helloWorld")      # "hello_world"
to_camel("Hello World")     # "helloWorld"
to_pascal("hello world")    # "HelloWorld"

# Text formatting
slugify("Hello World!")     # "hello-world"
truncate("Hello World", 5)  # "Hello..."

# Padding
pad("hi", 6, "0")           # "0hi00"
pad_left("hi", 5, "0")      # "000hi"
pad_right("hi", 5, "0")     # "hi000"

# Text operations
pluralize("cat", 2)         # "cats"
pluralize("sheep", 2)       # "sheep" (unchanged)
mask_text("12345678", 2, 2) # "12****78"
random_string(8)            # "aB3xY9kL"

# Text analysis
is_palindrome("racecar")    # True
word_count("Hello world")   # 2
```

## Installation

```bash
pip install sensible
```

Or from source:

```bash
git clone https://github.com/dominionthedev/exp.git
cd sensible/py
uv pip install -e .
```

## Usage

```python
from sensible.types import Result, Option
from sensible.chains import chain
from sensible.decorators import cached, validate_args, retry
from sensible.strings import slugify, pad, pluralize

# Use Result for safe error handling
res = Result.ok(5).map(lambda x: x * 2).unwrap()  # 10

# Use Option for optional values
opt = Option.some(42).filter(lambda x: x > 10).unwrap()  # 42

# Chain operations
result = chain([1, 2, 3]).then(sum).map(lambda x: x * 2).value()  # 12

# Caching
@cached(max_size=100)
def compute(x):
    return x * x

# Type validation
@validate_args(str, int)
def greet(name: str, age: int):
    return f"{name} ({age})"
```

## Tests

```bash
pytest tests/ -v
```

## License

MIT