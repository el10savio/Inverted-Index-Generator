package main

import (
	"fmt"
	"regexp"
	"strings"
)

func RemoveDuplicates(wordList []string) []string {
	keys := make(map[string]bool)
	uniqueWords := []string{}

	for _, entry := range wordList {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueWords = append(uniqueWords, entry)
		}
	}

	return uniqueWords
}

func Preprocessing(wordList []string) []string {
	ProcessedWordList := []string{}

	for _, word := range wordList {
		ProcessedWordList = append(ProcessedWordList, strings.ToLower(word))
	}

	return ProcessedWordList
}

func Tokenize(Doc string) []string {
	wordList := []string{}

	r := regexp.MustCompile("[^\\s]+")
	wordList = r.FindAllString(Doc, -1)

	wordList = Preprocessing(wordList)
	wordList = RemoveDuplicates(wordList)

	// fmt.Println(wordList)
	return wordList
}

func GenerateDocMap(token []string) map[string]bool {
	docMap := make(map[string]bool)

	for _, word := range token {
		if _, value := docMap[word]; !value {
			docMap[word] = true
		}
	}

	return docMap
}

func GenerateInvertedIndex(DocList []string) {
	// fmt.Println("DocList:", DocList)
	globalDocMap := make([]map[string]bool, 0)

	for _, Doc := range DocList {
		token := Tokenize(Doc)
		docMap := GenerateDocMap(token)
		globalDocMap = append(globalDocMap, docMap)
	}

	// fmt.Println("globalDocMap:", globalDocMap)
	var invertedIndex InvertedIndex

	for DocMapIndex, DocMap := range globalDocMap {
		for DocEntry, _ := range DocMap {
			var Item InvertedIndexItem
			Item.Term = DocEntry
			Item.Frequency = 1

			Item.PostingsList.value = DocMapIndex
			Item.PostingsList.next = nil

			invertedIndex.Items = append(invertedIndex.Items, Item)
		}
	}
	fmt.Println("invertedIndex:", invertedIndex)
}

func main() {
	DocList := []string{
		"new home sales top forecasts",
		"home sales rise in July",
		"increase in home sales in July",
		"July new home sales rise"}

	GenerateInvertedIndex(DocList)
}
