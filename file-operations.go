package main

import (
	"io/ioutil"
	"os"
)

func ReadContent(args Arguments) (error, []byte) {
	file, err := os.OpenFile(args["fileName"], os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		return ThrowError(err), nil
	}

	bytes, err := ioutil.ReadAll(file)
	file.Close()

	if err != nil {
		return ThrowError(err), nil
	}

	return nil, bytes
}
