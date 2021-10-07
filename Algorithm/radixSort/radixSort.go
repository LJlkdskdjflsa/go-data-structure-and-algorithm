package radixSort

import "fmt"

func SelectSortMax(arr []int) int {
	length := len(arr) //數組長度
	if length <= 1 {

		return arr[0]
	} else {
		max := arr[0]
		for i := 1; i < length; i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max
	}
}

func RadixSort(arr []int) []int {
	max := SelectSortMax(arr)              //尋找最大值
	for bit := 1; max/bit > 0; bit *= 10 { //等比區間
		//for bit := 1; max/bit > 0; bit += 10 {//等差區間
		arr = BitSort(arr, bit) //每次處理一個區間
		fmt.Println(arr)

	}
	return arr

}

func BitSort(arr []int, bit int) []int {
	length := len(arr)           //數組長度
	bitcounts := make([]int, 10) //統計長度 0,1,2,3,4,5,6,7,8,9
	for i := 0; i < length; i++ {
		num := (arr[i] / bit) % 10 //分層處理,當bit=100,二位數不參與排序
		bitcounts[num]++
	}
	fmt.Println(bitcounts)
	//0 1 2 3 4 5
	//1 1 2 3 4 5
	//1 1 1 4 4 5
	for i := 1; i < 10; i++ {
		bitcounts[i] += bitcounts[i-1] //疊加 計算位置
	}

	fmt.Println(bitcounts)
	tmp := make([]int, 10) //臨時數組
	for i := length - 1; i >= 0; i-- {
		num := (arr[i] / bit) % 10
		fmt.Println(bitcounts)
		//fmt.Println(len(arr))
		//fmt.Println(bitcounts[num])
		tmp[bitcounts[num]-1] = arr[i] //計算排序位置
		bitcounts[num]--               //往下循環
	}
	for i := 0; i < length; i++ {
		arr[i] = tmp[i] //保存數組
	}

	return arr
}
