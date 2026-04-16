from typing import Generic, TypeVar, Callable, Any, List, Optional, Iterable

T = TypeVar("T")
U = TypeVar("U")

class Chain(Generic[T]):
    """A fluent wrapper for composing operations on a value."""
    def __init__(self, value: T):
        self._value = value

    def then(self, func: Callable[[T], U]) -> "Chain[U]":
        """Apply a function to the current value and return a new Chain."""
        return Chain(func(self._value))

    def pipe(self, *funcs: Callable[[Any], Any]) -> "Chain[Any]":
        """Pipe the current value through a sequence of functions."""
        result = self._value
        for func in funcs:
            result = func(result)
        return Chain(result)

    def value(self) -> T:
        """Return the final value of the chain."""
        return self._value

    def map(self, func: Callable[[T], U]) -> "Chain[U]":
        """Apply a function to the value and return a new Chain."""
        return Chain(func(self._value))

    def filter(self, predicate: Callable[[T], bool]) -> "Chain[T]":
        """Filter the value using a predicate. Returns self if predicate is True, None if False."""
        if predicate(self._value):
            return self
        return Chain(None)

    def tap(self, func: Callable[[T], Any]) -> "Chain[T]":
        """Apply a function to the value for side effects, then return the value unchanged."""
        func(self._value)
        return self

    def if_else(
        self,
        condition: bool,
        then_func: Optional[Callable[[T], U]] = None,
        else_func: Optional[Callable[[T], U]] = None
    ) -> "Chain[U]":
        """Conditionally apply one of two functions based on a boolean condition."""
        if condition:
            return Chain(then_func(self._value)) if then_func else self
        return Chain(else_func(self._value)) if else_func else self

    def fold(self, func: Callable[[T, T], T]) -> "Chain[T]":
        """Fold/reduce the chain value by applying a function cumulatively."""
        # For now, just return self - can be extended for iterable values
        return self

    def collect(self) -> T:
        """Consume the chain and return the final value."""
        return self._value

    def if_none(self, default: T) -> "Chain[T]":
        """Return default value if current value is None."""
        if self._value is None:
            return Chain(default)
        return self

    def or_else(self, other: T) -> "Chain[T]":
        """Return other value if current value is None (alias for if_none)."""
        return self.if_none(other)

def chain(value: T) -> Chain[T]:
    """Start a new chain with the given value."""
    return Chain(value)

def chain_from_iterable(iterable: Iterable[T]) -> "Chain[List[T]]":
    """Start a new chain from an iterable."""
    return Chain(list(iterable))
