from dataclasses import dataclass
from typing import Generic, TypeVar, Optional, Any, Callable, Union, Type

T = TypeVar("T")
E = TypeVar("E")
U = TypeVar("U")

@dataclass(frozen=True)
class Result(Generic[T, E]):
    """A value that can be either a success or a failure."""
    _value: Optional[T] = None
    _error: Optional[E] = None
    _is_ok: bool = True

    @classmethod
    def ok(cls, value: T) -> "Result[T, Any]":
        return cls(_value=value, _is_ok=True)

    @classmethod
    def error(cls, error: E) -> "Result[Any, E]":
        return cls(_error=error, _is_ok=False)

    @property
    def is_ok(self) -> bool:
        return self._is_ok

    @property
    def is_error(self) -> bool:
        return not self._is_ok

    def unwrap(self) -> T:
        if not self._is_ok:
            raise ValueError(f"called unwrap on an error result: {self._error}")
        return self._value

    def unwrap_or(self, default: T) -> T:
        return self._value if self._is_ok else default

    def unwrap_or_else(self, func: Callable[[E], T]) -> T:
        """Apply the given function to the error and return the result if this is an error, otherwise return the value."""
        if self._is_ok:
            return self._value
        return func(self._error)

    def map(self, func: Callable[[T], U]) -> "Result[U, E]":
        """Apply a function to the value if it's ok, return the original error if not."""
        if self._is_ok:
            return Result.ok(func(self._value))
        return Result.error(self._error)

    def map_error(self, func: Callable[[E], Any]) -> "Result[T, Any]":
        """Apply a function to the error if it's an error, return the original value if ok."""
        if self._is_ok:
            return Result.ok(self._value)
        return Result.error(func(self._error))

    def and_then(self, func: Callable[[T], "Result[U, E]"]) -> "Result[U, E]":
        """Chain a function that returns another Result, skipping if already in error state."""
        if self._is_ok:
            return func(self._value)
        return Result.error(self._error)

    def error_value(self) -> E:
        if self._is_ok:
            raise ValueError("called error_value on a success result")
        return self._error

    def __str__(self) -> str:
        if self._is_ok:
            return f"Ok({self._value})"
        return f"Error({self._error})"

    def __repr__(self) -> str:
        return self.__str__()


@dataclass(frozen=True)
class Option(Generic[T]):
    """A value that can be either some value or nothing."""
    _value: Optional[T] = None
    _is_some: bool = True

    @classmethod
    def some(cls, value: T) -> "Option[T]":
        return cls(_value=value, _is_some=True)

    @classmethod
    def none(cls) -> "Option[Any]":
        return cls(_is_some=False)

    @property
    def is_some(self) -> bool:
        return self._is_some

    @property
    def is_none(self) -> bool:
        return not self._is_some

    def unwrap(self) -> T:
        if not self._is_some:
            raise ValueError("called unwrap on a None option")
        return self._value

    def unwrap_or(self, default: T) -> T:
        return self._value if self._is_some else default

    def unwrap_or_else(self, func: Callable[[], T]) -> T:
        """Apply the given function to get a default value if this is none."""
        if self._is_some:
            return self._value
        return func()

    def map(self, func: Callable[[T], U]) -> "Option[U]":
        """Apply a function to the value if it's some, return None if not."""
        if self._is_some:
            return Option.some(func(self._value))
        return Option.none()

    def map_or(self, default: U, func: Callable[[T], U]) -> U:
        """Apply a function to the value if it's some, otherwise return the default."""
        if self._is_some:
            return func(self._value)
        return default

    def filter(self, predicate: Callable[[T], bool]) -> "Option[T]":
        """Return None if the value doesn't match the predicate, otherwise return self."""
        if self._is_some and predicate(self._value):
            return self
        return Option.none()

    def and_then(self, func: Callable[[T], "Option[U]"]) -> "Option[U]":
        """Chain a function that returns another Option, skipping if already None."""
        if self._is_some:
            return func(self._value)
        return Option.none()

    def or_else(self, func: Callable[[], "Option[T]"]) -> "Option[T]":
        """Return self if some, otherwise apply the given function."""
        if self._is_some:
            return self
        return func()

    def __str__(self) -> str:
        if self._is_some:
            return f"Some({self._value})"
        return "None"

    def __repr__(self) -> str:
        return self.__str__()
