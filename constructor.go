package trier

import "fmt"

func Complete[T any](result T, err error) Try[T] {
	return Try[T]{
		Result: result,
		Error:  err,
	}
}

func Completef[T any](result T, format string, a ...any) Try[T] {
	return Complete(result, fmt.Errorf(format, a...))
}

func Success[T any](result T) Try[T] {
	return Complete(result, nil)
}

func Failure[T any](err error) Try[T] {
	var zero T
	return Complete(zero, err)
}

func Failuref[T any](format string, a ...any) Try[T] {
	return Failure[T](fmt.Errorf(format, a...))
}
