package main

import (
	"fmt"
	"strings"

	"example.com/m/bubbleSort"
	"example.com/m/heapSort"
	"example.com/m/insertSort"
	"example.com/m/mergeSort"
	"example.com/m/oddEvenSort"
	"example.com/m/quickSort"
	"example.com/m/radixSort"
	"example.com/m/selectSort"
	"example.com/m/shellSort"
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
func main7() {

	mySlice := []int{9, 8, 7, 6, 4, 31, 7, 5823, 12, 3}
	fmt.Println(quickSort.QuickSort(mySlice))
}

//odd even sort
func main8() {

	mySlice := []int{9, 8, 7, 6, 4, 31, 7, 5823, 12, 3}
	fmt.Println(oddEvenSort.OddEvenSort(mySlice))
}

//merge sort
func main9() {

	mySlice := []int{9, 8, 7, 6, 5, 4, 3, 7}
	fmt.Println(mergeSort.MergeSort(mySlice))
}

//shell sort
func main10() {

	mySlice := []int{9, 8, 7, 6, 5, 4, 3, 7}
	fmt.Println(shellSort.ShellSort(mySlice))
}

//radix sort
func main() {

	//mySlice := []int{9, 8, 7, 6, 5, 4, 3, 7, 58, 34, 73, 43, 53, 65, 7, 33, 53, 97, 43, 6, 77, 89, 66, 43, 68, 22, 15, 77, 43, 88, 57, 33, 79, 99, 22, 35, 84, 25, 74, 35, 75, 32}
	mySlice := []int{8, 7, 6, 5, 43, 88, 57, 33, 79, 99}
	//11 222 33
	//91 42  123
	fmt.Println(radixSort.RadixSort(mySlice))
}
