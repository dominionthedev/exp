package types

import (
	"errors"
	"fmt"
	"testing"
)

func TestResult(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		r := Ok(10)
		if !r.IsOk() {
			t.Error("expected IsOk to be true")
		}
		if r.Unwrap() != 10 {
			t.Errorf("expected value to be 10, got %d", r.Unwrap())
		}
	})

	t.Run("Error", func(t *testing.T) {
		err := errors.New("fail")
		r := Error[int](err)
		if !r.IsError() {
			t.Error("expected IsError to be true")
		}
		if r.UnwrapOr(0) != 0 {
			t.Errorf("expected UnwrapOr to be 0, got %d", r.UnwrapOr(0))
		}
	})

	t.Run("Map", func(t *testing.T) {
		r := Ok(10)
		r2 := MapResult(r, func(x int) string { return fmt.Sprintf("%d", x) })
		if r2.Unwrap() != "10" {
			t.Errorf("expected value to be \"10\", got %s", r2.Unwrap())
		}
	})

	t.Run("MapError", func(t *testing.T) {
		r := Error[int](errors.New("fail"))
		r2 := MapError(r, func(err error) error { return fmt.Errorf("wrapped: %w", err) })
		if r2.Error().Error() != "wrapped: fail" {
			t.Errorf("expected error to be \"wrapped: fail\", got %s", r2.Error().Error())
		}
	})

	t.Run("AndThen", func(t *testing.T) {
		r := Ok(10)
		r2 := AndThenResult(r, func(x int) Result[string] { return Ok(fmt.Sprintf("%d", x)) })
		if r2.Unwrap() != "10" {
			t.Errorf("expected value to be \"10\", got %s", r2.Unwrap())
		}

		r3 := AndThenResult(r, func(x int) Result[string] { return Error[string](errors.New("fail")) })
		if !r3.IsError() {
			t.Error("expected r3 to be an error")
		}
	})
}

func TestOption(t *testing.T) {
	t.Run("Some", func(t *testing.T) {
		o := Some("hello")
		if !o.IsSome() {
			t.Error("expected IsSome to be true")
		}
		if o.Unwrap() != "hello" {
			t.Errorf("expected value to be hello, got %s", o.Unwrap())
		}
	})

	t.Run("None", func(t *testing.T) {
		o := None[string]()
		if !o.IsNone() {
			t.Error("expected IsNone to be true")
		}
		if o.UnwrapOr("default") != "default" {
			t.Errorf("expected UnwrapOr to be default, got %s", o.UnwrapOr("default"))
		}
	})

	t.Run("Map", func(t *testing.T) {
		o := Some(10)
		o2 := OptionMap(o, func(x int) string { return fmt.Sprintf("%d", x) })
		if o2.Unwrap() != "10" {
			t.Errorf("expected value to be \"10\", got %s", o2.Unwrap())
		}
	})

	t.Run("Filter", func(t *testing.T) {
		o := Some(10)
		o2 := o.Filter(func(x int) bool { return x > 5 })
		if !o2.IsSome() {
			t.Error("expected o2 to be some")
		}

		o3 := o.Filter(func(x int) bool { return x > 15 })
		if !o3.IsNone() {
			t.Error("expected o3 to be none")
		}
	})

	t.Run("AndThen", func(t *testing.T) {
		o := Some(10)
		o2 := OptionAndThen(o, func(x int) Option[string] { return Some(fmt.Sprintf("%d", x)) })
		if o2.Unwrap() != "10" {
			t.Errorf("expected value to be \"10\", got %s", o2.Unwrap())
		}
	})

	t.Run("OrElse", func(t *testing.T) {
		o := None[int]()
		o2 := o.OrElse(func() Option[int] { return Some(10) })
		if o2.Unwrap() != 10 {
			t.Errorf("expected value to be 10, got %d", o2.Unwrap())
		}
	})
}
