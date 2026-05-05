package types

// Option represents a value that can be either some value or nothing.
type Option[T any] struct {
	value   T
	present bool
}

// Some creates a new Option with a value.
func Some[T any](val T) Option[T] {
	return Option[T]{value: val, present: true}
}

// None creates a new Option with no value.
func None[T any]() Option[T] {
	return Option[T]{present: false}
}

// IsSome returns true if the option has a value.
func (o Option[T]) IsSome() bool {
	return o.present
}

// IsNone returns true if the option has no value.
func (o Option[T]) IsNone() bool {
	return !o.present
}

// Unwrap returns the value if it's present, or panics if it's not.
func (o Option[T]) Unwrap() T {
	if !o.present {
		panic("called Unwrap on a None option")
	}
	return o.value
}

// UnwrapOr returns the value if it's present, or the given default value if it's not.
func (o Option[T]) UnwrapOr(def T) T {
	if !o.present {
		return def
	}
	return o.value
}

// UnwrapOrElse returns the value if it's present, or the result of the given function if it's not.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if !o.present {
		return f()
	}
	return o.value
}

// UnwrapOrDefault returns the value if it's present, or the zero value of T if it's not.
func (o Option[T]) UnwrapOrDefault() T {
	if !o.present {
		var zero T
		return zero
	}
	return o.value
}

// IsSomeAnd returns true if the option is Some and the value matches the predicate.
func (o Option[T]) IsSomeAnd(predicate func(T) bool) bool {
	return o.present && predicate(o.value)
}

// Value returns the value and a boolean (standard Go map pattern).
func (o Option[T]) Value() (T, bool) {
	return o.value, o.present
}

// OptionMap applies a function to the value if it's present, returning a new Option.
func OptionMap[T, U any](o Option[T], f func(T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return Some(f(o.value))
}

// Filter returns None if the value doesn't match the predicate, otherwise returns the original Option.
func (o Option[T]) Filter(predicate func(T) bool) Option[T] {
	if o.IsSome() && predicate(o.value) {
		return o
	}
	return None[T]()
}

// OptionAndThen chains a function that returns another Option.
func OptionAndThen[T, U any](o Option[T], f func(T) Option[U]) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return f(o.value)
}

// OrElse returns the original Option if it's some, otherwise applies the given function.
func (o Option[T]) OrElse(f func() Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}
	return f()
}
