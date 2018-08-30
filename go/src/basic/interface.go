package basic
import "fmt"
import "math"

/*
	Interfaces are the way to achieve polymorphism in golang.
	Just forget about oops for a moment and think in interface manner.

	Flexibility provided by interfaces comes from the fact that single type can implement many interfaces

	or the same interface can be satisfied by many types
*/
type geometry interface {
	area() float64
	perim() float64

}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}


func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(d geometry) (float64, float64) {
	return d.area(), d.perim()
}

func interfaceConversion(){
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}



// blank interface
func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func NullInterface(f interface{}){
	fmt.Println(f)
	fmt.Printf("(%v, %T)\n", f, f)
}

func TestInterface(){
	var g geometry
	g  = rect{10, 12}
	fmt.Println(g.perim(), g.area())

	c := circle{radius: 5}
	c_area, c_per := measure(c)
	fmt.Println(c_area, c_per)


	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
	PrintAll(vals)
	NullInterface(g)
	interfaceConversion()

}