package main

import (
	"fmt"
	"sort"
)

func main() {
	EvenArray := []int{9, 7, 2, 5, 3, 6}
	OddArray := []int{2, 5, 11, 6, 8, 9, 10}
	fmt.Println(median(EvenArray))
	fmt.Println(median(OddArray))
}

// Returns the median of the array
func median(i []int) float64 {
	sort.Ints(i)

	if len(i)%2 == 1 {
		return float64(i[len(i)/2])
	}

	return float64(i[len(i)/2]+i[len(i)/2-1]) / 2
}
