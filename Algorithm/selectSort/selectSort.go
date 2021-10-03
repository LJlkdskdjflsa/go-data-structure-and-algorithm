package selectSort

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
func SelectSort(arr []int) []int {
	length := len(arr) //數組長度
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			//只剩一個元素不需要挑選
			min := i // 索引標記
			for j := i + 1; j < length; j++ {
				//每次選出一個極小值
				if arr[min] < arr[j] {
					min = j //極小值索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //最小值不等於當前值,就交換

			}
		}
	}

	return arr
}
