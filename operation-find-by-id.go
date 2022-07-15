package main

import (
	"encoding/json"
	"io"
)

func findById(args Arguments, writer io.Writer) error {
	if len(ItemId) == 0 {
		return ThrowError(ErrorMustDefineTheIdFlag)
	}

	err, item, _, _ := findElementById(args, ItemId)

	if err != nil {
		return ThrowError(err)
	}

	if item == nil {
		writer.Write([]byte(""))
		return nil
	}

	bytes, err := json.Marshal(item)
	if err != nil {
		return ThrowError(err)
	}

	writer.Write(bytes)

	return nil
}
