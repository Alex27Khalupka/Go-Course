package main

import "fmt"

// max gets a slice of strings and
// returns string - the longest word in the slice.
func max(StrArray []string) string {
	var max = ""
	for _, element := range StrArray {
		if len(element) > len(max) {
			max = element
		}
	}
	return max
}

func main() {
	StrSliceOne := []string{"one", "two", "three"}
	StrSliceTwo := []string{"one", "two"}
	StrSliceThree := []string{"one", "two", "three", "seventeen", "OK"}

	fmt.Println(max(StrSliceOne))
	fmt.Println(max(StrSliceTwo))
	fmt.Println(max(StrSliceThree))
}
