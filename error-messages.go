package main

import "errors"

var (
	ErrorMustDefineTheOperationFlag = errors.New("-operation flag has to be specified")
	ErrorMustDefineTheFileNameFlag  = errors.New("-fileName flag has to be specified")
	ErrorMustDefineTheItemFlag      = errors.New("-item flag has to be specified")
	ErrorMustDefineTheIdFlag        = errors.New("-id flag has to be specified")
)
