package basic

import (
	"fmt"
	"strconv"
)

func client(i chan int){
	var n = []int{1, 4, 9, 8}
	for _, ind := range n{
		fmt.Println("sending value" + strconv.Itoa(ind))
		i <-ind
	}
	close(i)
}


func server(cli chan int, res chan int){
	for data := range cli{
		fmt.Println("receiving value" +strconv.Itoa(data))
		res <- data
	}
	close(res)

}

func TestConcurrency() {
	cli := make(chan int)
	ser := make(chan int)
	go client(cli)
	go server(cli, ser)
	for d := range ser {
		fmt.Println("Getting")
		fmt.Println(d)
	}

}