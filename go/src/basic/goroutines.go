package basic

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func test(){
	fmt.Println("name is this")
}

func TestGoRoutines() {
	go say("world")
	test()
}