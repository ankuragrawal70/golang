package basic

import "fmt"

func PassPointer(){
	
}


func TestPointer() {
	i, j := 42, 2701

	var p *int
	p = &i         // point to i
	fmt.Println(*p) // read i through the pointer
	
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// pointer initialized without var variable
	k := 50
	x := &k
	fmt.Println(k, *(&k), *x)
	*x = 100
	fmt.Println(k, *(&k), *x)
}
