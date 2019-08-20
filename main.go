package main

import (
	processing "./processing"
)

func main() {
	DocList := []string{
		"new home sales top forecasts",
		"home sales rise in July",
		"increase in home sales in July",
		"July new home sales rise"}

	index := processing.GenerateInvertedIndex(DocList)

	processing.Find(index, "Sales")
	processing.Find(index, "June")
}
