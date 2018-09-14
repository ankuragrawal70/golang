package sorting

func Merge(arr[] int, l, m, h int)  {
	tmp_array_1 := make([]int, m-l+1)
	tmp_array_2 := make([]int, h-m)
	copy(tmp_array_1, arr[l:m+1])
	copy(tmp_array_2, arr[m+1:h+1])
	i, j, k:= 0, 0, l

	for (i<len(tmp_array_1) && j<len(tmp_array_2)) {
		if tmp_array_1[i] < tmp_array_2[j] {
			arr[k] = tmp_array_1[i]
			i += 1
		} else if tmp_array_1[i] == tmp_array_2[j] {
			arr[k] = tmp_array_1[i]
			i += 1
			j += 1
		} else {
			arr[k] = tmp_array_2[j]
			j += 1
		}
		k += 1
	}
		if i<len(tmp_array_1){
			for i<len(tmp_array_1){
				arr[k] = tmp_array_1[i]
				i++
				k++
			}
		}
		if j<len(tmp_array_2){
			for j<len(tmp_array_2){
				arr[k] = tmp_array_2[j]
				j++
				k++
			}
		}
}

func MergeSort(arr[] int, l, r int){
	if l <  r {
		mid := (l + r) / 2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)
	}
	return
}

//func main()  {
//	const cap = 10000000
//	var arr [cap]int
//	var initial_array = arr[:]
//	for i:=0; i<cap; i++{
//		initial_array[i] = rand.Intn(cap)
//	}
//	c_time := time.Now()
//	// fmt.Println(initial_array)
//	fmt.Println(c_time)
//	MergeSort(initial_array, 0, len(initial_array)-1)
//	e_time := time.Now()
//	fmt.Println(e_time.Sub(c_time))
//}


