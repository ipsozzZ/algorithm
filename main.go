package main

import (
	"algorithm/array"
	"fmt"
)

/* ===========================================
	算法测试入口，algorithm下子包不再实现单元测试代码
*/
func main() {
	fmt.Printf("ni hao algorithm")
	var numSlice = []int{7, 8, 30, 22, 44, 1, 7}

	/*// 选择排序
	array.HelloSelection()
	var selectionArr = array.SelectionSort(numSlice)
	fmt.Printf("\nselectionArr => %v", selectionArr)

	// 冒泡排序
	array.HelloBubble()
	var bubbleArr = array.BubbleSort(numSlice)
	fmt.Printf("\nbubbleArr => %v", bubbleArr)*/

	// 快速排序
	array.HelloQuick()
	var quickArr = array.QuickSort(numSlice)
	fmt.Printf("\nquickArr => %v", quickArr)
}
