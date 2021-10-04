package heapSort

import "fmt"

func HeapSortMax(arr []int, length int) []int {
	//length := len(arr) //數組長度
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1         //深度
		for i := depth; i >= 0; i-- { //循環所有的三節點
			topmax := i                                                //假定最大的在i的位置
			leftchild := topmax*i + 1                                  //左節點
			rightchild := topmax*i + 2                                 //右節點
			if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越界
				topmax = arr[leftchild] //如果左邊大,紀錄最大
			}

			if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
				topmax = arr[rightchild] //如果右邊大,紀錄最大
			}
			if topmax != i { //確保i的值最大
				arr[i], arr[topmax] = arr[topmax], arr[i]
			}
		}
		return arr
	}
}
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastmesslen := length - 1 //每次擷取一段
		HeapSortMax(arr, lastmesslen)
		fmt.Println(arr)
		if i < length {
			arr[0], arr[lastmesslen] = arr[lastmesslen], arr[0]
		}
		fmt.Println("ex:", arr)
	}
	return arr
}
