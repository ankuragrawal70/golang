package sorting



func InsertionSort(arr []int){
	for initial_v:=1; initial_v<len(arr); initial_v++{
		//fmt.Println(arr)
		v := arr[initial_v]
		if v >= arr[initial_v-1]{
			continue
		}
		tmp := initial_v - 1 
		//fmt.Println(tmp)
		for arr[tmp]>=v && tmp>0{
			arr[tmp+1] = arr[tmp]
			tmp -=1  
		}
		if arr[tmp] <= v{
			arr[tmp+1] = v
		} else{
			arr[tmp] = v
		}
		//fmt.Println(tmp)
		
		
		
	}
}

//func main(){
//
//	var arr [cap]int
//	var initial_array = arr[:]
//	for i:=0; i<cap; i++{
//		initial_array[i] = rand.Intn(cap)
//	}
//	c_time := time.Now()
//	// fmt.Println(initial_array)
//	fmt.Println(c_time)
//	InsertionSort(initial_array)
//	e_time := time.Now()
//	fmt.Println(e_time.Sub(c_time))
//	//fmt.Println("Final value is", initial_array)
//}
