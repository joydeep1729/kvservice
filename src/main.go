package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		fmt.Printf("the error is: %v\n", e)
	}
}

func main() {
	var err error
	KeyValMap, err = loadData("Corpus.CSV")
	check(err)
	fmt.Println(KeyValMap)
	server(":8080")
}
