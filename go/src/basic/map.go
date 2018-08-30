package basic
import "fmt"

// map is passed as call by reference
func passMap(t map[string]int) {
	t["0"] = 10
	fmt.Println("modified map is", t["0"])
}

func MapCheck(){
	var a = make(map[string]int);
	a["0"] = 1;
	passMap(a)
	fmt.Println("successful map", a["0"])	
	
	// always check ok here as 0 is a valid value
	k, ok1 := a["1"]

	if res, ok := a["1"]; ok {
		fmt.Println(res)
	}
	
	key, ok := a["0"]
	fmt.Println(key, ok, k, ok1)

}