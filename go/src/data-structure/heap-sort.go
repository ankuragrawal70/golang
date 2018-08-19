package main
import "fmt"
import "math/rand"
import "time"
// import "reflect"


const cap = 10000000

func heapifyDown(i int, arr []int){
	// fmt.Println("in heapify", arr)
	if i >=len(arr){
		return
	}
	index := i
	max := arr[i]
	l := 2*i+1
	r := 2*i+2
	// if arr[l] == nil && arr[r] == nil{
	// 	return
	// }
	// fmt.Println(reflect.TypeOf(l))
	if l < len(arr) && (arr[l] > max){
		
		max = arr[l]
		index = l
		// fmt.Println("hh", max, arr[r])
	} 
	if r < len(arr) && (arr[r] > max){

		max = arr[r]
		index = r
	}
	// fmt.Println(i, index)
	if index != i{
		arr[i], arr[index] = arr[index], arr[i]
		heapifyDown(index, arr)
	}
	return
	
}

func createHeap(arr []int){
	start := (len(arr)-1)/2
	// fmt.Println("start is", start)
	for i:=start; i>=0; i--{
		heapifyDown(i, arr)
	}
}

func heapSort(arr []int){
	createHeap(arr)
	// fmt.Println("after heap creation", arr)
	l := len(arr)
	
	for i:=l; i>0; i--{
		arr[0], arr[i-1] = arr[i-1], arr[0]
		// fmt.Println("after top re hepify", arr)
		heapifyDown(0, arr[0:i-1])
		// fmt.Println("after top removing", arr)
	}
}

func main(){
	
	var initial_array = make([]int, cap)
	// cap := 100000
	// var initial_array [cap]int
	// var initial_array = arr[:]
	for i:=0; i<cap; i++{
		initial_array[i] = rand.Intn(cap)
	}
	c_time := time.Now()
	fmt.Println(c_time)
	// fmt.Println("Initial array is", initial_array[0:len(initial_array)])
	heapSort(initial_array)
	e_time := time.Now()
	fmt.Println(e_time.Sub(c_time))
	// fmt.Println("Final value is", initial_array)
}