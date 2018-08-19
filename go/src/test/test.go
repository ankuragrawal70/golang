package test

import "fmt"
import "time"
// import "math"
func T() {

	c_time := time.Now()
	// fmt.Println(initial_array)
	fmt.Println(c_time)
	fmt.Printf("vvalue is %v", testFun(10, 16666611, 1000000000))
	e_time := time.Now()
	fmt.Println("time is", e_time.Sub(c_time))
    
}

func fun(x float64) float64 {
	// math.Pow(x, 2) is very slow here
	return x*x - x
}

func testFun(first float64, second float64, times int) float64 {
    var i = 0
    var result float64 = 0
    var dx float64
    dx = (second - first) / float64(times)
    for ; i < times; i++ {
        result += fun(first + float64(i)*dx)
	}
    return result * dx
}