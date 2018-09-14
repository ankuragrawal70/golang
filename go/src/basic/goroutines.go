package basic

import (
	"fmt"
	"time"
	"sync"
)

func testChannels(){
	messages := make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		messages <- 1
	}()
	go func() {
		time.Sleep(time.Second * 2)
		messages <- 2
	}()
	go func() {
		time.Sleep(time.Second * 1)
		messages <- 3
	}()
	func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()
	//time.Sleep(time.Second * 4)
}



var w sync.WaitGroup

func say(s string, w *sync.WaitGroup) {
	defer w.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	//w.Done()
}

func test(w *sync.WaitGroup){
	defer w.Done()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("test success")
	}
	//w.Done()
}

func TestGoRoutines() {
	//w.Add(2)
	//go say("world", &w)
	//go test(&w)
	//w.Wait()
	testChannels()
}