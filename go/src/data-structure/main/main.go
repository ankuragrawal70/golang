package main

import (
	"time"
	"fmt"
	"data-structure/searching"
	"math/rand"
	"data-structure/sorting"
)


type MS func(arr []int, l, h int)
type GSort func(arr []int)

func evaluateTime(initial_array []int, f MS){
	c_time := time.Now()
	// fmt.Println(initial_array)
	fmt.Println(c_time)
	f(initial_array, 0, len(initial_array)-1)
	e_time := time.Now()
	fmt.Println(e_time.Sub(c_time))
}

func evaluateTimeSort(initial_array []int, f GSort){
	c_time := time.Now()
	// fmt.Println(initial_array)
	fmt.Println(c_time)
	f(initial_array)
	e_time := time.Now()
	fmt.Println(e_time.Sub(c_time))
}


func main()  {
	const cap = 1000000
	var arr [cap]int
	var initial_array = arr[:]
	for i:=0; i<cap; i++{
		initial_array[i] = rand.Intn(cap)
		//initial_array[i] = i+1
	}
	//fmt.Println(initial_array[0:100])

	fmt.Println("time for merge sort for ", cap, "elements")
	evaluateTime(initial_array, sorting.MergeSort)
	fmt.Println("ntime for insertion sort for ", cap, "elements")
	evaluateTimeSort(initial_array[:], sorting.HeapSort)
	evaluateTimeSort(initial_array[:], sorting.InsertionSort)
	searching.BinarySearch([]int{2, 15, 18, 20, 25, 28, 44}, 28)
}