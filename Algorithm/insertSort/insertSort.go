package insertSort

import "fmt"

func InsertSort(arr []int) []int {
	length := len(arr) //數組長度
	if length <= 1 {
		return arr
	} else {
		for i := 1; i < length; i++ {
			//跳過第一個
			backup := arr[i] //備份要插入的數
			j := i - 1       //往回找插入位置(從上個位置循環)
			for j >= 0 && backup < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = backup
			fmt.Println(arr)
		}
	}
	return arr
}
