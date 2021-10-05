package quickSort

func QuickSort(arr []int) []int {
	length := len(arr) //數組長度
	if length <= 1 {
		return arr
	} else {
		splitdata := arr[0]       //基準
		low := make([]int, 0, 0)  //存儲較小的
		high := make([]int, 0, 0) //存儲較大的
		mid := make([]int, 0, 0)  //相等
		mid = append(mid, splitdata)
		for i := 1; i < length; i++ { //循環跳過第一個數
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] >= splitdata {

				high = append(high, arr[i])
			} else {

				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) //切割遞迴
		myarray := append(append(low, mid...), high...)
		return myarray
	}
}
