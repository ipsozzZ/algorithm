package array

import "fmt"

func HelloQuick() {
	fmt.Printf("\nni hao quick_sort")
}

/* ========================================
	快速排序实现接口
	@Copyright   ...
	@Author      ipso
*/
func QuickSort(arr []int) []int {
	arrLen := len(arr)
	doSort(arr, 0, arrLen-1)
	return arr
}


/*--------------------------- 内部函数 -----------------------*/

///
func doSort(arr []int, low int, high int) {

	if low < high {
		var index = getIndex(arr, low, high)
		doSort(arr, low, index-1)
		doSort(arr, index+1, high)
	}
}

/// 获取下一轮排序起始下标
func getIndex(arr []int, low int, high int) int {
	var tmp = arr[low]
	for low < high  {
		// 从尾部开始向前寻找第一个小于中间值的元素，找到结束循环，否则high减1
		for low < high && arr[high] >= tmp {
			high--
		}
		arr[low] = arr[high] // 找到小于中间值的high并将其与low交换
		// 从low开始向后寻找第一个大于中间值的元素，找到结束循环，否则low加1
		for low < high && arr[low] <= tmp {
			low++
		}
		arr[high] = arr[low] // 找到将
	}
	arr[low] = tmp
	return low
}