package basic
import "fmt"
/*
default channels are unbuffered, meaning that 
they will only accept sends (chan <-) if there 
is a corresponding receive (<- chan) ready to 
receive the sent value. Buffered channels accept 
a limited number of values without a 
corresponding receiver for those values.
*/
func sum(arr []int, c chan int){
	sum := 0
	for _, v:= range arr{
		sum += v
	}
	c <- sum
}

func bufferChannel(){
	// buffered channel does not wait for receiver untill the buffer is full
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func bufferChannelDeadlock(){
	//reducing channel capacity and assigning more values cause deadlock
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func RangeClose(){
	// closed by sender to intimate that no other data is sent
	fmt.Println("range and close started")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}


func fibonacciTest(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}


func sTest(c, quit chan int){
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

func testSelectChannel(){
	var c = make(chan int)
	var quit = make(chan int)
	go sTest(c, quit)
	fibonacciTest(c, quit)
}


func TestChannel(){
	// arr := []int{10, 5, 12, 18, 18}
	// c := make(chan int)

	// go sum(arr[0:len(arr)/2], c)
	// go sum(arr[len(arr)/2:], c)
	// x, y := <-c, <-c
	// fmt.Println(x, y, x+y)

	// bufferChannel()
	// RangeClose()
	// bufferChannelDeadlock()
	testSelectChannel()
}