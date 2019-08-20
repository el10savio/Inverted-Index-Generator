package invertedIndex

// InvertedIndexItem contains the term followed by the
// number of times it has appeared across all documents
// and an array of documents it is persent in
type InvertedIndexItem struct {
	Term            string
	Frequency       int
	DocumentListing []int
}

// InvertedIndex contains a hash map to easily check if the
// term is present and an array of InvertedIndexItem
type InvertedIndex struct {
	HashMap map[string]*InvertedIndexItem
	Items   []InvertedIndexItem
}

// FindItem returns the position of a given
// Item in an Inverted Index
func (invertedIndex *InvertedIndex) FindItem(Term string) int {
	for index, item := range invertedIndex.Items {
		if item.Term == Term {
			return index
		}
	}
	panic("Not Found")
}

// AddItem works by first checking if a given term is already present
// in the inverse index or not by checking the hashmap. If it is
// present it updates the Items by increasing the frequency and
// adding the document it is found in. If it is not present it
// adds it to the hash map and adds it to the items list
func (invertedIndex *InvertedIndex) AddItem(Term string, Document int) {
	if invertedIndex.HashMap[Term] != nil {
		// log.Println("Index item", Term, "already exists :: updating existing item")

		FoundItemPosition := invertedIndex.FindItem(Term)

		invertedIndex.Items[FoundItemPosition].Frequency++
		invertedIndex.Items[FoundItemPosition].DocumentListing = append(invertedIndex.Items[FoundItemPosition].DocumentListing, Document)
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

// CreateInvertedIndex initializes an
// empty Inverted Index
func CreateInvertedIndex() *InvertedIndex {
	invertedIndex := &InvertedIndex{
		HashMap: make(map[string]*InvertedIndexItem),
		Items:   []InvertedIndexItem{},
	}
	return invertedIndex
}
