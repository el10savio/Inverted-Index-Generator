package invertedIndex

type LinkedList struct {
	value int
	next  *LinkedList
}

type InvertedIndexItem struct {
	Term         string
	Frequency    int
	PostingsList LinkedList
}

type InvertedIndex struct {
	Items []InvertedIndexItem
}
