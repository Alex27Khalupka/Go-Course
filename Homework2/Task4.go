package main

import (
	"fmt"
	"sort"
)

// printSorted gets map(key - integer, value - string),
// prints map values sorted in order of increasing keys.
func printSorted(_map map[int]string) {
	keys := make([]int, 0, len(_map))
	for key, _ := range _map {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Print(key, ": ", _map[key], "; ")
	}
	fmt.Println()
}

func main() {
	map1 := map[int]string{2: "a", 0: "b", 1: "c"}
	map2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}

	printSorted(map1)
	printSorted(map2)
}
