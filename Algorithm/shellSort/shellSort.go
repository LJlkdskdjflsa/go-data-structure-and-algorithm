package shellSort

import "fmt"

func ShellSortStep(arr []int, start int, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap {
		//below from insert sort
		backup := arr[i] //備份要插入的數
		j := i - gap     //往回找插入位置(從上個位置循環)
		for j >= 0 && backup < arr[j] {
			arr[j+gap], arr[j] = arr[j], arr[j+gap] //互換
			j -= gap
		}
		//arr[j+gap] = backup //插入
		fmt.Println(arr)
	}
}

func ShellSort(arr []int) []int {
	length := len(arr) //數組長度
	if length <= 1 {
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ {
				ShellSortStep(arr, i, gap)

			}
			gap /= 2
		}
	}
	return arr
}
