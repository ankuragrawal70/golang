package basic
import "fmt"

func Loop(){
	fmt.Print("successfully ommited")
	for i:=0; i<10; i++{
		fmt.Print(i, " ")
	}

	// like while
	var sum =10
	for sum<20{
		sum+=2
		fmt.Print(sum, " ")
	}
	fmt.Println("")
	//range loop on iterator
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v:= range list{
		fmt.Println(k, v)
	}

}