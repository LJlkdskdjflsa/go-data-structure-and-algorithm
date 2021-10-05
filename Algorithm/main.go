package main

import (
	"fmt"
	"strings"

	"example.com/m/bubbleSort"
	"example.com/m/heapSort"
	"example.com/m/insertSort"
	"example.com/m/quickSort"
	"example.com/m/selectSort"
)

// select sort int
func main1() {
	mySlice := []int{1, 2, 3, 4, 6, 3, 4, 5, 2, 9, 8, 6, 3, 0}
	fmt.Println(selectSort.SelectSortMax(mySlice))
	fmt.Println(selectSort.SelectSort(mySlice))
}

//compare strings
func main2() {
	//相等 0,小於 -1,大於 1
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("c", "b"))
	fmt.Println(strings.Compare("a2", "a1"))
	fmt.Println(strings.Compare("a1", "a2"))
	//1.3版本後適用
	fmt.Println("a123" > "a456")
	fmt.Println("a" > "b")
	fmt.Println("b" > "a")
}

func main3() {
	mySlice := []string{"a", "b", "c", "d", "e", "asdf", "asdc", "bgr"}
	//fmt.Println(selectSort.SelectSortMax(mySlice))
	fmt.Println(selectSort.SelectSortString(mySlice))
}

// insert sort int
func main4() {
	mySlice := []int{9, 8, 7, 6, 5, 4, 7, 6, 5}
	fmt.Println(insertSort.InsertSort(mySlice))

}

// bubble sort
func main5() {
	mySlice := []int{9, 8, 7, 6, 5, 4, 7, 6, 5}
	fmt.Println(bubbleSort.BubbleSort(mySlice))

}

//heap sort
func main6() {

	mySlice := []int{9, 8, 7, 6, 5, 4, 7, 6, 5}
	fmt.Println(heapSort.HeapSort(mySlice))
}

//quick sort
func main() {

	mySlice := []int{9, 8, 7, 6, 4, 31, 7, 5823, 12, 3}
	fmt.Println(quickSort.QuickSort(mySlice))
}
