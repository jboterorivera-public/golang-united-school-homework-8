package main

import "encoding/json"

func findElementById(args Arguments, itemId string) (error, *ItemDTO, int, []ItemDTO) {
	err, bytes := ReadContent(args)
	if err != nil {
		return ThrowError(err), nil, -1, nil
	}

	var items []ItemDTO

	if len(bytes) > 0 {
		err = json.Unmarshal(bytes, &items)

		if err != nil {
			return ThrowError(err), nil, -1, nil
		}
	}

	for index, item := range items {
		if item.Id == itemId {
			return nil, &item, index, items
		}
	}

	return nil, nil, -1, items
}
