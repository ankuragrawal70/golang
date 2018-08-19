package basic
import "fmt"


// pass through references

type Vertex struct {
	X int
	Y int
}

func passStrunct(v *Vertex){
	v.X = 100
	fmt.Println(*v)
}

func TestStruct(){
	x := Vertex{20, 11}
	x.X = 40
	fmt.Println(x)


	var p *Vertex
	p = &x

	// both ways are possible
	(*p).X = 199
	p.Y = 50
	fmt.Println(*p)

	//passing 
	passStrunct(&x)
	fmt.Println(x)
}