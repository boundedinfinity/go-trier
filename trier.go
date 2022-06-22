package trier

import "fmt"

type Try[T any] struct {
	Result T
	Err    error
}

func (t Try[T]) Failure() bool {
	return t.Err != nil
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
	return FailureR(zero, err)
}

func FailureR[T any](result T, err error) Try[T] {
	return Try[T]{
		Result: result,
		Err:    err,
	}
}

func Failuref[T any](format string, a ...any) Try[T] {
	var zero T
	return FailureRf(zero, format, a...)
}

func FailureRf[T any](result T, format string, a ...any) Try[T] {
	return Try[T]{
		Result: result,
		Err:    fmt.Errorf(format, a...),
	}
}
