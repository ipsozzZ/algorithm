package array

import (
	"fmt"
)

func HelloBubble() {
	fmt.Printf("\nni hao bubble_sort")
}

/* =====================================
	冒泡排序
	@Author ipso
	@Title  冒泡排序
*/
func BubbleSort(arr []int) []int {
	var arrLen = len(arr)
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp     := arr[j]
				arr[j]   = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
