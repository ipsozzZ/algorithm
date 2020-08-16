package array

import "fmt"

func HelloSelection() {
	fmt.Printf("\nni hao selection_sort")
}

/* =====================================
	选择排序
	@Author ipso
	@Title  选择排序
*/
func SelectionSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
