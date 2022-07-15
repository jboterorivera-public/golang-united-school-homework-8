package main

import "io"

var operationsList = map[string]func(args Arguments, writer io.Writer) error{
	"add":      add,
	"findById": findById,
	"list":     list,
	"remove":   remove,
}

func GetFunction(operation string) (bool, func(args Arguments, writer io.Writer) error) {
	f, exists := operationsList[operation]

	return exists, f
}
