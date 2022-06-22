package errorer

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Errorf(format string, a ...any) Error {
	return Err(fmt.Errorf(format, a...))
}

func Err(err error) Error {
	return Error{
		err: err,
	}
}

func None() Error {
	return Error{
		err: nil,
	}
}

type Error struct {
	err error
}

func (e Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}

	return ""
}

func (e Error) MarshalJSON() ([]byte, error) {
	if e.err != nil {
		return json.Marshal(e.Error())
	}

	return json.Marshal(nil)
}

func (e *Error) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e.err = errors.New(s)

	return nil
}
