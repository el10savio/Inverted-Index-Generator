package processing

import (
	"fmt"
	"regexp"
	"strings"

	iIndex "../invertedIndex"
)

// RemoveDuplicates filters out all duplicate
// words from each document
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

// Preprocessing converts each word to lowercase
// TODO: Clean up each word for symbols
func Preprocessing(wordList []string) []string {
	ProcessedWordList := []string{}

	// Convert each string to lowercase from
	// wordList and add to ProcessedWordList
	for _, word := range wordList {
		ProcessedWordList = append(ProcessedWordList, strings.ToLower(word))
	}

	return ProcessedWordList
}

// Tokenize gets the individual words from each
// document and generates a wordlist
func Tokenize(Doc string) []string {
	wordList := []string{}

	// The following regexp finds indivdual
	//  words in a sentence
	r := regexp.MustCompile("[^\\s]+")
	wordList = r.FindAllString(Doc, -1)

	wordList = Preprocessing(wordList)
	wordList = RemoveDuplicates(wordList)

	return wordList
}

// GenerateDocMap creates a hash map of
// each word in the document
func GenerateDocMap(token []string) map[string]bool {
	docMap := make(map[string]bool)

	for _, word := range token {
		if _, value := docMap[word]; !value {
			docMap[word] = true
		}
	}

	return docMap
}

// GenerateInvertedIndex for each document list
// gets each word as a token, processes it and
// generates a hash map for each document
// using them it then generates the
// inverted index of all words
func GenerateInvertedIndex(DocList []string) iIndex.InvertedIndex {
	globalDocMap := make([]map[string]bool, 0)

	for _, Doc := range DocList {
		token := Tokenize(Doc)
		docMap := GenerateDocMap(token)
		globalDocMap = append(globalDocMap, docMap)
	}

	// Create an empty inverted index
	invertedIndex := iIndex.CreateInvertedIndex()

	// Using the generated hash maps add
	// each word to the inverted index
	for DocMapIndex, DocMap := range globalDocMap {
		for DocEntry, _ := range DocMap {
			invertedIndex.AddItem(DocEntry, DocMapIndex)
		}
	}
	return *invertedIndex
}

// Find for a given inverted index and search term
// checks if the term exists and then
// ouputs the documents the
// term is in
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
