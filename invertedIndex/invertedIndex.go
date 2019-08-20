package invertedIndex

type InvertedIndexItem struct {
	Term            string
	Frequency       int
	DocumentListing []int
}

type InvertedIndex struct {
	HashMap map[string]*InvertedIndexItem
	Items   []InvertedIndexItem
}

func (invertedIndex *InvertedIndex) FindItem(Term string) int {
	for index, item := range invertedIndex.Items {
		if item.Term == Term {
			return index
		}
	}
	panic("Not Found")
}

func (invertedIndex *InvertedIndex) AddItem(Term string, Document int) {
	if invertedIndex.HashMap[Term] != nil {
		// log.Println("Index item", Term, "already exists :: updating existing item")

		InvertedIndexItemPosition := invertedIndex.FindItem(Term)

		invertedIndex.Items[InvertedIndexItemPosition].Frequency++
		invertedIndex.Items[InvertedIndexItemPosition].DocumentListing = append(invertedIndex.Items[InvertedIndexItemPosition].DocumentListing, Document)
	} else {
		// log.Println("Index item", Term, " does not exist :: creating new item")

		invertedIndexItem := &InvertedIndexItem{
			Term:            Term,
			Frequency:       1,
			DocumentListing: []int{Document},
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
