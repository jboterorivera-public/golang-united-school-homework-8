package main

func ValidateFlags(args Arguments) error {
	ItemId = args["id"]
	Item = args["item"]
	Operation = args["operation"]
	FileName = args["fileName"]

	if len(Operation) == 0 {
		return ThrowError(ErrorMustDefineTheOperationFlag)
	}

	if len(FileName) == 0 {
		return ThrowError(ErrorMustDefineTheFileNameFlag)
	}

	return nil
}
