package main

import (
	"encoding/json"
	"io"
	"os"
)

func add(args Arguments, writer io.Writer) error {
	if len(Item) == 0 {
		return ThrowError(ErrorMustDefineTheItemFlag)
	}

	var itemToCreate ItemDTO
	err := json.Unmarshal([]byte(Item), &itemToCreate)
	if err != nil {
		return ThrowError(err)
	}

	err, itemInFile, _, itemsInFile := findElementById(args, itemToCreate.Id)

	if err != nil {
		return ThrowError(err)
	}

	if itemInFile != nil {
		writer.Write([]byte("Item with id " + itemToCreate.Id + " already exists"))
		return nil
	}

	result := append(itemsInFile, itemToCreate)

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
