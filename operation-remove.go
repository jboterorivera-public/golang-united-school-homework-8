package main

import (
	"encoding/json"
	"io"
	"os"
)

func remove(args Arguments, writer io.Writer) error {
	if len(ItemId) == 0 {
		return ThrowError(ErrorMustDefineTheIdFlag)
	}

	err, itemInFile, itemIndex, itemsInFile := findElementById(args, ItemId)

	if err != nil {
		return ThrowError(err)
	}

	if itemInFile == nil {
		writer.Write([]byte("Item with id " + ItemId + " not found"))
		return nil
	}

	result := append(itemsInFile[:itemIndex], itemsInFile[itemIndex+1:]...)

	bytes, err := json.Marshal(result)
	if err != nil {
		return ThrowError(err)
	}

	fileName := args["fileName"]

	if err := os.Truncate(fileName, 0); err != nil {
		return ThrowError(err)
	}

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	file.Write(bytes)
	file.Close()

	return nil
}
