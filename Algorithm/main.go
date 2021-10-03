package main

import (
	"fmt"

	"example.com/m/selectSort"
)

// select sort
func main() {
	mySlice := []int{1, 2, 3, 4, 6, 3, 4, 5, 2, 9, 8, 6, 3, 0}
	fmt.Println(selectSort.SelectSortMax(mySlice))
	fmt.Println(selectSort.SelectSort(mySlice))
}
