package trier

import (
	"fmt"

	"github.com/boundedinfinity/go-trier/errorer"
)

type Try[T any] struct {
	Result T     `json:"result"`
	Error  error `json:"error"`
}

func (t Try[T]) Failure() bool {
	return t.Error != nil
}

func (t Try[T]) Success() bool {
	return !t.Failure()
}

func Success[T any](result T) Try[T] {
	return Try[T]{
		Result: result,
	}
}

func Failure[T any](err error) Try[T] {
	var zero T
	return FailureWithResult(zero, err)
}

func FailureWithResult[T any](result T, err error) Try[T] {
	err2 := errorer.Err(err)
	return Try[T]{
		Result: result,
		Error:  &err2,
	}
}

func Failuref[T any](format string, a ...any) Try[T] {
	return Failure[T](fmt.Errorf(format, a...))
}

func FailurefWithResult[T any](result T, format string, a ...any) Try[T] {
	return FailureWithResult(result, fmt.Errorf(format, a...))
}
