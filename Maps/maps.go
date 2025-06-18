package main

import (
	"fmt"
	"strings"
)

//! Create a map to store product prices. Add and remove products.
//! Count frequency of each word in a sentence using map.

func main() {
	store := map[string]int{"a": 1, "b": 2, "c": 3}
	sentence := "One Two Three Two One"

	//* For Count frequency
	words := strings.Fields(sentence)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	for word, count := range frequency {
		fmt.Println(word, "=>", count)
	}

	//* for  Add and remove products
	store["d"] = 4
	fmt.Println("Store: ", store)
	delete(store, "b")
	fmt.Println("After: ", store)
}
