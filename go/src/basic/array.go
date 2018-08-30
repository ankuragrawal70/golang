package basic
import "fmt"

// passed as call by value
func passArray(arr [3]int) {
	arr[1] = 10
	fmt.Println("Array is", arr)
}

func ArrayCheck(){
	var te = 50
	var a [3]int;
	a[0] = 1;
	passArray(a)
	for x, y:= range a{
		fmt.Println("x and y", x, y)
	}

	for i:=0; i<len(a); i++{
		fmt.Println(i)
	}
	fmt.Println("successful", a, te)	
}


