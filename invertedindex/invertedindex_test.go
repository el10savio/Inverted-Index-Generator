package invertedindex

import (
	"reflect"
	"testing"
)

var wordListTest = []string{
	"new", "HOme", "sales", "top", "forecasts",
	"home", "sales", "rise", "in", "July",
	"increase", "in", "home", "SALES", "in",
	"July", "new", "home", "sales", "rise", "July",
}

func TestPreprocessing(t *testing.T) {
	wordList := wordListTest

	expectedList := []string{
		"new", "home", "sales", "top", "forecasts",
		"home", "sales", "rise", "in", "july",
		"increase", "in", "home", "sales", "in",
		"july", "new", "home", "sales", "rise", "july",
	}

	actualList := Preprocessing(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestPreprocessing_NoWordList(t *testing.T) {
	wordList := make([]string, 0)

	expectedList := make([]string, 0)

	actualList := Preprocessing(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	wordList := Preprocessing(wordListTest)

	expectedList := []string{
		"new", "home", "sales", "top", "forecasts",
		"rise", "in", "july",
		"increase",
	}

	actualList := RemoveDuplicates(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestRemoveDuplicates_NoWordList(t *testing.T) {
	wordList := make([]string, 0)

	expectedList := make([]string, 0)

	actualList := RemoveDuplicates(wordList)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestTokenize(t *testing.T) {
	doc := "new home sales top forecasts NEW"

	expectedList := []string{
		"new", "home", "sales", "top", "forecasts",
	}

	actualList := Tokenize(doc)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}

func TestTokenize_NoDoc(t *testing.T) {
	var doc string

	expectedList := []string{}

	actualList := Tokenize(doc)

	if !reflect.DeepEqual(expectedList, actualList) {
		t.Fatalf("\nExpected:%v \nGot:%v", expectedList, actualList)
	}
}
