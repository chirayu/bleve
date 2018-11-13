package main

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/index/upsidedown"

	_ "net/http/pprof"
)


func main() {

	index, _ := bleve.Open("example.bleve")


	internalIndex, _, err := index.Advanced()
	if err != nil {
		panic("1")
		return
	}
	internalIndexReader, err := internalIndex.Reader()
	if err != nil {
		panic("1")
		return
	}

	var rv []interface{}
	rowChan := internalIndexReader.DumpAll()

	for row := range rowChan {
		fmt.Printf("%T\n", row)
		fmt.Println(row)
		switch row := row.(type) {
		case error:
			panic("1")
			return
		case upsidedown.UpsideDownCouchRow:
			tmp := struct {
				Key []byte `json:"key"`
				Val []byte `json:"val"`
			}{
				Key: row.Key(),
				Val: row.Value(),
			}
			rv = append(rv, tmp)
		}
	}
	err = internalIndexReader.Close()
	if err != nil {
		panic("1")
		return
	}

	// fmt.Println(rv)
}
