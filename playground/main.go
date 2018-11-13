package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
	_ "net/http/pprof"
)


func main() {

	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.New("example.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := struct {
		Name string
	}{
		Name: "text",
	}
		
	// index some data
	index.Index("id", data)

	barData := struct {
		BarName string
		FooName string
	}{
		BarName: "bar",
		FooName: "foo",
	}

	// index some data
	index.Index("id1", barData)


	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)

}
