package bubbleSort

func BubbleSort(arr []int) []int {
	length := len(arr) //數組長度
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ {
			need_sort := false //確認是否已排序
			//只剩一個就不需要冒泡了
			for j := 0; j < length-1-i; j++ {

				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					need_sort = true
				}
			}
			if !need_sort {
				return arr
			}
		}
		return arr
	}
}
