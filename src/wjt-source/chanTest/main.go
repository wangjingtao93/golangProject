package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	go func() {
		c <- 1
		time.Sleep(time.Second)
		c <- 2
		time.Sleep(time.Second)
		close(c)
	}()

	for i := 0; i < 6; i++ {
		j, ok := <-c
		fmt.Printf("receive: %d, status: %t\n", j, ok)
	}
}