package searching

import "fmt"

func SearchTmp(arr []int, l, h int, n int) int{

	if l<h{
		mid := l + (h-l)/2
		if arr[mid] == n{
			return mid
		}else if arr[mid] > n{
				h = mid-1

				}else{
					l=mid+1
		}
		return SearchTmp(arr, l, h, n)
		}else if arr[l] == n{
			return l
	}
	return -1
	}


func BinarySearch(arr []int, n int){
	v := SearchTmp(arr, 0, len(arr)-1, n)
	fmt.Println(v)
}

