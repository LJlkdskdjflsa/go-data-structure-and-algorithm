package mergeSort

import "fmt"

func merge(leftArr []int, rightArr []int) []int {
	leftIndex := 0
	rightIndex := 0
	resultArr := []int{}
	for leftIndex < len(leftArr) && rightIndex < len(rightArr) {
		if leftArr[leftIndex] < rightArr[rightIndex] {
			resultArr = append(resultArr, leftArr[leftIndex])
			leftIndex++
		} else if leftArr[leftIndex] > rightArr[rightIndex] {

			resultArr = append(resultArr, rightArr[rightIndex])
			rightIndex++
		} else {
			//相等

			resultArr = append(resultArr, leftArr[leftIndex])
			leftIndex++
			resultArr = append(resultArr, rightArr[rightIndex])
			rightIndex++
		}
		//假設一組陣列填完了,令一組就直接全下

		/*
			if leftIndex < len(leftArr) {
				resultArr = append(resultArr, leftArr[leftIndex:]...)
			}

			if rightIndex < len(rightArr) {
				resultArr = append(resultArr, rightArr[rightIndex:]...)
			}
		*/
	}
	for leftIndex < len(leftArr) {
		resultArr = append(resultArr, leftArr[leftIndex])
		leftIndex++
	}

	for rightIndex < len(rightArr) {
		resultArr = append(resultArr, rightArr[rightIndex])
		rightIndex++
	}
	fmt.Println(resultArr)
	return resultArr
}

func MergeSort(arr []int) []int {
	length := len(arr) //數組長度
	if length <= 1 {
		return arr

	} else {
		mid := length / 2
		leftArr := MergeSort(arr[:mid])
		rightArr := MergeSort(arr[mid:])
		//fmt.Println(leftArr)
		//fmt.Println(rightArr)
		return merge(leftArr, rightArr)
	}
}
