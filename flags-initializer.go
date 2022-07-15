package main

import "flag"

var ItemId string
var Item string
var Operation string
var FileName string

func InitFlags() {
	flag.StringVar(&ItemId, "id", "", "Unique Item Id in order to process certain operations")
	flag.StringVar(&Item, "item", "", "Item in json format in order to process certain operations")
	flag.StringVar(&Operation, "operation", "", "Performs some allowed operations like list, add, remove and findById")
	flag.StringVar(&FileName, "fileName", "", "Filename for performing the requested operations")

	flag.Parse()
}
