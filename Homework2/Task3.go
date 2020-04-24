package main

import "fmt"

// reverse function gets []int64 slice and returns []int64 slice that is
// a reversed copy of received slice
func reverse(slice []int64) []int64 {
	reversedSlice := make([]int64, len(slice), cap(slice))
	if len(slice) == 0 {
		return reversedSlice
	}
	for i := 0; i < len(slice)/2+1; i++ {
		// position opposite to position i in the slice
		opp := len(slice) - i - 1
		reversedSlice[i] = slice[opp]
		reversedSlice[opp] = slice[i]
	}
	return reversedSlice
}

func main() {
	SliceOne := []int64{1, 2, 5, 15}
	SliceTwo := []int64{1, 2, 3, 4, 5}
	PalindromeSlice := []int64{5, 6, 7, 6, 5}
	OneElementSlice := []int64{15}
	var EmptySlice []int64

	fmt.Println(SliceOne, ", reversed:", reverse(SliceOne))
	fmt.Println(SliceTwo, ", reversed:", reverse(SliceTwo))
	fmt.Println(PalindromeSlice, ", reversed:", reverse(PalindromeSlice))
	fmt.Println(OneElementSlice, ", reversed:", reverse(OneElementSlice))
	fmt.Println(EmptySlice, ", reversed:", reverse(EmptySlice))
}
