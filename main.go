package main

import (
	"fmt"
	"strconv"

	"github.com/blevesearch/bleve"
)

var index bleve.Index

func main() {
	createIndex()
	addItemsToIndex()
	getAll()
	searchFullText()
	deleteItemWithID()
	searchFullTextByCategory()
}

func createIndex() {
	var err error
	index, err = bleve.Open("news")
	if err != nil {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New("news", mapping)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func deleteItemWithID() {
	getItemWithTitle("01Titlea example")
	fmt.Println("\ndeleting title '01Titlea example' item with id")
	index.Delete("News item 1")
	getItemWithTitle("01Titlea example")
}

func getItemWithTitle(title string) {
	fmt.Println("\ngetting item with title '" + title + "'")
	query := bleve.NewMatchQuery(title)

	query.SetField("Title")
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults.Hits.Len(), "search results")
	printSearchResults(searchResults)
}

func searchFullText() {
	fmt.Println("\ngetting items which contain 'Politics' or 'Science' text in any of the text fields")
	query := bleve.NewMatchQuery("example Science")
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	search.Size = 10
	search.SortBy([]string{"Title"})
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSearchResults(searchResults)
}

func printSearchResults(searchResults *bleve.SearchResult) {
	for i := range searchResults.Hits {
		fmt.Println(searchResults.Hits[i].ID, searchResults.Hits[i].Index,
			searchResults.Hits[i].Fields["Name"],
			searchResults.Hits[i].Fields["Category"],
			searchResults.Hits[i].Fields["Title"])
	}
}

func searchFullTextByCategory() {
	fmt.Println("\ngetting 10 first items which have 'Science' category sorted by title")
	query := bleve.NewMatchQuery("Science")
	query.SetField("Category")
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	search.Size = 10
	search.SortBy([]string{"Title"})
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSearchResults(searchResults)
}

func getAll() {
	fmt.Println("\ngetting 10 first of all titles in index sorted by title")
	query := bleve.NewMatchAllQuery()
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	search.Size = 10
	search.SortBy([]string{"Title"})
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSearchResults(searchResults)
}

func addItemsToIndex() {
	index.Index("News item 1", data{Name: "BBC", Title: "01Titlea example", Category: "Politics"})
	index.Index("News item 2", data{Name: "Financial Times", Title: "02_Title_b", Category: "Economy"})
	index.Index("News item 3", data{Name: "BBC", Title: "03_Title_c", Category: "Entertainment"})
	index.Index("News item 4", data{Name: "Telegraph", Title: "04_Title_d", Category: "Economy"})
	index.Index("News item 5", data{Name: "BBC", Title: "05_Title_e", Category: "Science"})
	index.Index("News item 6", data{Name: "Guardian", Title: "06_Title_f", Category: "Politics"})
	index.Index("News item 7", data{Name: "BBC", Title: "07_Title_g", Category: "Politics"})
	index.Index("News item 8", data{Name: "BBC", Title: "08_Title_h", Category: "Science"})
	index.Index("News item 9", data{Name: "Daily Mail", Title: "09_Title_i", Category: "Economy"})
	index.Index("News item 10", data{Name: "BBC", Title: "10_Title_j", Category: "Politics"})
	index.Index("News item 11", data{Name: "Mashable", Title: "11_Title_k", Category: "Science"})
	index.Index("News item 12", data{Name: "BBC", Title: "12_Title_l", Category: "Travel"})
	index.Index("News item 13", data{Name: "BBC", Title: "13_Title_m", Category: "Politics"})

	for i := 14; i < 100; i++ {
		index.Index("News item "+strconv.Itoa(i), data{Name: "Test", Title: "Title_test", Category: "Politics"})
	}
}

type data struct {
	Name     string
	Title    string
	Category string
}
