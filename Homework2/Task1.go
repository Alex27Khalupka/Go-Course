package main

import "fmt"

// Returns average value of the array
func average(array [6]int) float64 {
	var sum = 0
	for _, element := range array {
		sum += element
	}
	return float64(sum) / float64(len(array))
}

func main() {
	var FirstArray = [6]int{1, 2, 3, 4, 5, 6}
	var SecondArray = [6]int{8, 9, 3, 4, 5, 6}
	var ThirdArray = [6]int{1, 1, 1, 1, 1, 1}

	fmt.Println(average(FirstArray))
	fmt.Println(average(SecondArray))
	fmt.Println(average(ThirdArray))
}
