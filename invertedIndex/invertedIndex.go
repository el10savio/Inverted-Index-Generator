package invertedIndex

import "fmt"

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

func (invertedIndex *InvertedIndex) AddItem(Item InvertedIndexItem) {
	if invertedIndex.HashMap[Item.Term] != nil {
		fmt.Println("Index item already exists")
	} else {
		fmt.Println("Index item does not exist")
	}
}
