package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{9, 7, 2, 5, 3, 6}
	fmt.Println(median(array))
}

func median(i []int) float64{
	sort.Ints(i)
	var med float64

	if len(i) % 2 == 1{
		med = float64(i[len(i) / 2])
		return med

	} else {
		med = float64(i[len(i) / 2] + i[len(i) / 2 - 1]) / 2
		return med
	}
}


