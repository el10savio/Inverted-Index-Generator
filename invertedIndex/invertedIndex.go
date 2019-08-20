package invertedIndex

import (
	"fmt"
)

type LinkedList struct {
	Value int
	Next  *LinkedList
}

type InvertedIndexItem struct {
	Term         string
	Frequency    int
	PostingsList LinkedList
}

type InvertedIndex struct {
	HashMap map[string]*InvertedIndexItem
	Items   []InvertedIndexItem
}

func (invertedIndex *InvertedIndex) AddItem(Term string, Document int) {
	if invertedIndex.HashMap[Term] != nil {
		fmt.Println("Index item already exists :: updating existing item")
	} else {
		fmt.Println("Index item does not exist :: creating new item")

		var _PostingsList LinkedList
		_PostingsList.Value = Document
		_PostingsList.Next = nil

		invertedIndexItem := &InvertedIndexItem{
			Term:         Term,
			Frequency:    1,
			PostingsList: _PostingsList,
		}

		invertedIndex.HashMap[Term] = invertedIndexItem
		invertedIndex.Items = append(invertedIndex.Items, *invertedIndexItem)

	}
}

func CreateInvertedIndex() *InvertedIndex {
	invertedIndex := &InvertedIndex{
		HashMap: make(map[string]*InvertedIndexItem),
		Items:   []InvertedIndexItem{},
	}
	return invertedIndex
}
