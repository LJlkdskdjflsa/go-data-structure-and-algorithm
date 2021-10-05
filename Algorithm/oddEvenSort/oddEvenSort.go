package oddEvenSort

//奇偶排序
//可以多線程
func OddEvenSort(arr []int) []int {
	//tmp := 0          //臨時數據
	isSorted := false //是否還需要排序(機數位,偶數位都不需要排序的有序)

	for isSorted == false {
		isSorted = true
		//基數位跟相鄰的偶數位互換,換到不需要換為止
		//基數位

		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {

				arr[i], arr[i+1] = arr[i+1], arr[i]

				isSorted = false
			}
		}
		//偶數位
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {

				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
	}
	return arr
}
