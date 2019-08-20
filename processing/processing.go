package processing

import (
	"fmt"
	"regexp"
	"strings"

	iIndex "../invertedIndex"
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

func GenerateInvertedIndex(DocList []string) iIndex.InvertedIndex {
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
	return *invertedIndex
}

func Find(index iIndex.InvertedIndex, searchTerm string) {
	Term := strings.ToLower(searchTerm)

	if index.HashMap[Term] != nil {
		itemPosition := index.FindItem(Term)
		item := index.Items[itemPosition]

		fmt.Println("Found:", searchTerm, "in documents:", item.DocumentListing)
	} else {
		fmt.Println("Not Found:", searchTerm)
	}
}
