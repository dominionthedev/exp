package chains

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

// Since Go doesn't support generic methods on structs with new type parameters,
// we use top-level functions for 'Then' to allow type transformation.
// However, for same-type transformations, we can have a method.

// Map applies a function that returns the same type.
func (c Chain[T]) Map(f func(T) T) Chain[T] {
	return Chain[T]{value: f(c.value)}
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
