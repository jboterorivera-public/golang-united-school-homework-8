package main

import (
	"io"
)

func list(args Arguments, writer io.Writer) error {
	err, bytes := ReadContent(args)

	if err != nil {
		return ThrowError(err)
	}

	writer.Write(bytes)

	return nil
}
