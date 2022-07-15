package main

import "fmt"

type customErrorStruct struct {
	Message string
	Err     error
}

func (c customErrorStruct) Error() string {
	return c.Message
}

func (c customErrorStruct) Unwrap() error {
	return c.Err
}

func ThrowError(e error) (c customErrorStruct) {
	c = customErrorStruct{Message: e.Error()}
	c.Err = fmt.Errorf("Error %w", e)

	return c
}
