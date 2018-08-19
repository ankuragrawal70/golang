package basic
import "fmt"

// slices are call by references

func passSlice(arr []int) {
	// arr[1] = 10
	// fmt.Println("slice is", arr)
}

func SliceCheck(){
	// var a = [3]int{1, 19, 19};
	// fmt.Println("Array is", a)
	// var sl []int = a[1:3]
	// fmt.Println("slice is", sl, len(sl), cap(sl))
	// sl[0] = 18
	
	// fmt.Println("slice after change is", sl)
	// fmt.Println("array is", a)


	var x = []int{4, 5}
	fmt.Println("new slice is", x, len(x), cap(x))
	// passed as reference
	passSlice(x)
	fmt.Println("Appending in slice")
	v := append(x, 24)
	n := append(v, 40)
	fmt.Println("new value", v, len(v), cap(v))
	fmt.Println("original slice is", x, len(x), cap(x))
	fmt.Println ("value is", n, len(x), cap(n))
	fmt.Println("original slice is", x, len(x), cap(x))


	// // nil slice
	// var s []int
	// fmt.Println(s, len(s), cap(s))
	// if s == nil {
	// 	fmt.Println("nil!")
	// }


	// //allocates a slice of length 5
	// b := make([]int, 5)
	// fmt.Println("value is", b)

	// //allocates a slice of lenght 0 and capacity 5
	// c := make([]int, 0, 5)
	// fmt.Println("value is", c)




	// append on nil slice
	// slice capacity is eep on increasing in multiplication factor(1, 2, 4, 8, 16, 32)
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3)
	printSlice(s)
		
	s = append(s, 20, 30)
	printSlice(s)

	s = append(s, 2, 3)
	printSlice(s)
		
	s = append(s, 2, 3, 4)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	s = append(s, 2, 3, 4, 8, 10, 10, 11, 12, 12, 18, 19)
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}