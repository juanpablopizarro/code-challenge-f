package main

import (
	"fmt"

	"github.com/juanpablopizarro/code-challenge-f/parser"
)

func main() {
	m, err := parser.Unmarshal([]byte("11AB398765UJ1A050200N23"))
	if err != nil {
		panic(err)
	}

	for k, v := range m {
		fmt.Printf("\tindex: %v, value: %v\n", k, v)
	}
}
