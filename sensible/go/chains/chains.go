package chains

import "reflect"

// Chain represents a fluent wrapper for composing operations on a value.
type Chain[T any] struct {
	value T
}

// New starts a new chain with the given value.
func New[T any](val T) Chain[T] {
	return Chain[T]{value: val}
}

// Then applies a function to the current value and returns a new Chain.
func Then[T, U any](c Chain[T], func_ func(T) U) Chain[U] {
	return Chain[U]{value: func_(c.value)}
}

// Value returns the final value of the chain.
func (c Chain[T]) Value() T {
	return c.value
}

// Collect returns the final value of the chain (alias for Value).
func (c Chain[T]) Collect() T {
	return c.value
}

// Since Go doesn't support generic methods on structs with new type parameters,
// we use top-level functions for 'Then' to allow type transformation.
// However, for same-type transformations, we can have a method.

// Map applies a function that returns the same type.
func (c Chain[T]) Map(f func(T) T) Chain[T] {
	return Chain[T]{value: f(c.value)}
}

// Filter returns a chain with the zero value of T if the predicate is false.
func (c Chain[T]) Filter(predicate func(T) bool) Chain[T] {
	if predicate(c.value) {
		return c
	}
	var zero T
	return Chain[T]{value: zero}
}

// Tap applies a function to the value for side effects, then returns the chain unchanged.
func (c Chain[T]) Tap(f func(T)) Chain[T] {
	f(c.value)
	return c
}

// Pipe pipes the current value through a sequence of functions that return the same type.
func (c Chain[T]) Pipe(funcs ...func(T) T) Chain[T] {
	val := c.value
	for _, f := range funcs {
		val = f(val)
	}
	return Chain[T]{value: val}
}

// IfNone returns a new Chain with the defaultValue if the current value is nil.
// This only applies to types that can be nil (pointers, interfaces, maps, slices, etc.).
func (c Chain[T]) IfNone(defaultValue T) Chain[T] {
	if isNil(c.value) {
		return Chain[T]{value: defaultValue}
	}
	return c
}

func isNil(i any) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	}
	return false
}

// OrElse returns a new Chain with the defaultValue if the current value is nil (alias for IfNone).
func (c Chain[T]) OrElse(defaultValue T) Chain[T] {
	return c.IfNone(defaultValue)
}
