package index

import (
	"github.com/jmmal/simple-search/analyser"
	"encoding/json"
	"fmt"
)

// Doc is a basic document used to index
type Doc struct {
	Id int32
	Name string
	Comment string
}

// DocumentFrequency is a map with key DocumentID and value frequency
type DocumentFrequency = map[int32]int32

// InvertedIndex is the engine used to query the data
type InvertedIndex struct {
	TermsIndex map[string]DocumentFrequency
	documents map[int32]*Doc
}

// NewIndex returns a new inverted index instance
func NewIndex() *InvertedIndex {
	return &InvertedIndex{
		TermsIndex: make(map[string]DocumentFrequency),
		documents: make(map[int32]*Doc),
	}
}

func (i *InvertedIndex) addTermToIndex(term string, docID, freq int32) {
	if _, ok := i.TermsIndex[term]; !ok {
		i.TermsIndex[term] = make(DocumentFrequency)
	}

	i.TermsIndex[term][docID] = freq
}

func (i *InvertedIndex) storeDocument(doc *Doc) {
	i.documents[doc.Id] = doc
}

// GetDocument returns the document in the index with the given id, if it exists
func (i *InvertedIndex) GetDocument(id int32) (doc *Doc, err error) {
	doc, ok := i.documents[id]
	
	if !ok {
		return doc, fmt.Errorf("Document does not exist")
	}

	return doc, nil
}

// Query searchers for documents matching the given query
func (i *InvertedIndex) Query(query string) (docs []*Doc, err error) {
	searchTokens := analyser.Tokenize(query)

	// Map is used to store the accumlative frequency of term matches for each
	// document
	documentTotalFrequency := make(map[int32]int32)

	for _, token := range searchTokens {
		termFreq, ok := i.TermsIndex[token]

		if !ok {
			continue
		}

		for id, freqForTerm := range termFreq {
			freq := documentTotalFrequency[id]

			documentTotalFrequency[id] = freq + freqForTerm
		}
	}

	fmt.Printf("%v\n", documentTotalFrequency)

	results := []*Doc{}

	// TODO: Sort the results by their frequency
	for docID := range documentTotalFrequency {
		doc, _ := i.GetDocument(docID)
		
		results = append(results, doc)
	}

	return results, nil
}

// IndexDocument will add the given document to the index
func (i *InvertedIndex) IndexDocument(doc *Doc) {
	var terms []string
	
	terms = append(terms, analyser.Tokenize(doc.Name)...)
	terms = append(terms, analyser.Tokenize(doc.Comment)...)

	for _, term := range terms {
		i.addTermToIndex(term, doc.Id, 1)
	}

	i.storeDocument(doc)
}

// PrintDocs pretty prints to the list of documents currently in the index
func (i *InvertedIndex) PrintDocs() {
	docs := []Doc{}

	for _, doc := range i.documents {
		docs = append(docs, *doc)
	}

	pretty, err := json.MarshalIndent(docs, "", "  ")

	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(string(pretty))
}

// PrintIndex will pretty print the current state of the index
func (i *InvertedIndex) PrintIndex() {
	for key, value := range i.TermsIndex {
		fmt.Printf("Term: '%s'\n", key)
		fmt.Println("Documents that contain the term:")

		docs := []Doc{}

		for id := range value {
			doc, err := i.GetDocument(id)

			if err != nil {
				continue
			}

			docs = append(docs, *doc)
		}
		
		prettyDocs, err := json.MarshalIndent(docs, "", "  ")

		if err != nil {
			fmt.Printf("Error formatting docs: %s\n", err)
		}

		fmt.Printf("%s\n\n\n", string(prettyDocs))
	}
}

// PrettyDocs return a PrettyString representation of a list of docs
func PrettyDocs(docs []*Doc) string {
	derefedDocs := []Doc{}

	for _, doc := range docs {
		derefedDocs = append(derefedDocs, *doc)
	}

	pretty, err := json.MarshalIndent(derefedDocs, "", "  ")

	if err != nil {
		fmt.Println("error: ", err)
	}

	return string(pretty)
}