package basic

// function with receiver
//The receiver appears in its own argument list between the func keyword and the method name.

import (
	"fmt"
	"math"
)

type VertexNew struct {
	X, Y float64
}

func (v VertexNew) Abs() float64 {
	v.X = 50
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func CheckReceiver(){
	v_new := VertexNew{10, 5}
	fmt.Println(v_new.Abs())
	fmt.Println(v_new)
}

