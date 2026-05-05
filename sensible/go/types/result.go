package types

import "fmt"

// Result represents a value that can be either a success or a failure.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a new successful Result.
func Ok[T any](val T) Result[T] {
	return Result[T]{value: val, err: nil}
}

// Error creates a new failed Result.
func Error[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// IsOk returns true if the result is a success.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsError returns true if the result is a failure.
func (r Result[T]) IsError() bool {
	return r.err != nil
}

// Unwrap returns the value if it's a success, or panics if it's a failure.
func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(fmt.Sprintf("called Unwrap on an error result: %v", r.err))
	}
	return r.value
}

// UnwrapOr returns the value if it's a success, or the given default value if it's a failure.
func (r Result[T]) UnwrapOr(def T) T {
	if r.err != nil {
		return def
	}
	return r.value
}

// UnwrapOrElse returns the value if it's a success, or the result of the given function applied to the error.
func (r Result[T]) UnwrapOrElse(f func(error) T) T {
	if r.err != nil {
		return f(r.err)
	}
	return r.value
}

// UnwrapOrDefault returns the value if it's a success, or the zero value of T if it's a failure.
func (r Result[T]) UnwrapOrDefault() T {
	if r.err != nil {
		var zero T
		return zero
	}
	return r.value
}

// IsOkAnd returns true if the result is Ok and the value matches the predicate.
func (r Result[T]) IsOkAnd(predicate func(T) bool) bool {
	return r.err == nil && predicate(r.value)
}

// IsErrorAnd returns true if the result is Error and the error matches the predicate.
func (r Result[T]) IsErrorAnd(predicate func(error) bool) bool {
	return r.err != nil && predicate(r.err)
}

// Error returns the error if it's a failure, or nil if it's a success.
func (r Result[T]) Error() error {
	return r.err
}

// Value returns the value and the error (standard Go return pattern).
func (r Result[T]) Value() (T, error) {
	return r.value, r.err
}

// MapResult applies a function to the value if it's a success, returning a new Result.
func MapResult[T, U any](r Result[T], f func(T) U) Result[U] {
	if r.IsError() {
		return Error[U](r.err)
	}
	return Ok(f(r.value))
}

// MapError applies a function to the error if it's a failure, returning a new Result.
func MapError[T any](r Result[T], f func(error) error) Result[T] {
	if r.IsOk() {
		return r
	}
	return Error[T](f(r.err))
}

// AndThenResult chains a function that returns another Result.
func AndThenResult[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.IsError() {
		return Error[U](r.err)
	}
	return f(r.value)
}
