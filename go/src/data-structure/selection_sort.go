package main
import "fmt"
import "math/rand"
import "time"
const cap = 100000

func SelectionSort(arr [cap]int){
	for i, e:= range arr{
		min := e
		flag := 0
		for j:=i+1; j<len(arr); j++{
			if arr[j] < min{
				min = arr[j]
				flag = j
			}
		}
		if flag > 0{
		arr[i], arr[flag] = arr[flag], arr[i]
		flag = 0
		}
	}
}

func main(){
	
	// var initial_array = make([]int, cap)
	// cap := 100000
	var initial_array [cap]int
	// var initial_array = arr[:]
	for i:=0; i<cap; i++{
		initial_array[i] = rand.Intn(cap)
	}
	c_time := time.Now()
	fmt.Println(c_time)
	SelectionSort(initial_array)
	e_time := time.Now()
	fmt.Println(e_time.Sub(c_time))
	// fmt.Println("Final value is", initial_array)
}


