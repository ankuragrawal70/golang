package sorting

import "fmt"

func partition(arr []int, l, h int) int {
	if h-l+1 < 2{
		return l
	}
	start := l
	end := h-1
	pivot := arr[h]
	//pivot_pos := l
	for start < end{
		for arr[start] < pivot{
			start ++
		}
		for arr[end] >= pivot{
			end--
		}
		if start < end {
			fmt.Println(end)
			arr[start], arr[end] = arr[end], arr[start]
		}
	}
	arr[start], arr[h] = arr[h], arr[start]
	return start
}


func QuickSort(arr []int, l, h int) {
	if l < h {
		index := partition(arr, l, h)
		QuickSort(arr, l, index-1)
		QuickSort(arr, index+1, h)
	}
}
func QuickSortMain(arr []int)  {
	//arr := []int{10, 8, 9, 1, 19, 7, 18, 28, 28, 22}
	QuickSort(arr, 0, len(arr)-1)
	//fmt.Println(arr)
}