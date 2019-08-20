package main

import (
	"fmt"
	"regexp"
	"strings"

	iIndex "./invertedIndex"
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
	globalDocMap := make([]map[string]bool, 0)

	for _, Doc := range DocList {
		token := Tokenize(Doc)
		docMap := GenerateDocMap(token)
		globalDocMap = append(globalDocMap, docMap)
	}

	invertedIndex := iIndex.CreateInvertedIndex()

	for DocMapIndex, DocMap := range globalDocMap {
		for DocEntry, _ := range DocMap {
			invertedIndex.AddItem(DocEntry, DocMapIndex)
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
