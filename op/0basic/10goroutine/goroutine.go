package main
import (
	"fmt"
	//"time"
)
/*
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}

// 通道 ch := make(chan int)
func sum(s []int, c chan int){
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
func main() {
	s := []int{1,2,3,4,5}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	//fmt.Println(s[:len(s)/2])
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
*/

func fibonacci(n int, c chan int) {
	x,y := 0,1
	for i := 0; i < n; i++ {
		c <- x
		x,y = y, x+y
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}