package main

import (
	"errors"
	"io"
	"os"
)

type Arguments map[string]string

func Perform(args Arguments, writer io.Writer) error {
	flagsError := ValidateFlags(args)

	if flagsError != nil {
		return flagsError
	}

	existsOperation, functionToRun := GetFunction(Operation)

	if !existsOperation {
		return ThrowError(errors.New("Operation " + Operation + " not allowed!"))
	}

	functionError := functionToRun(args, writer)

	if functionError != nil {
		return functionError
	}

	return nil
}

func main() {
	InitFlags()

	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	args := Arguments{
		"id":        ItemId,
		"operation": Operation,
		"item":      Item,
		"fileName":  FileName,
	}

	return args
}
