/*
How to use chanel
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	defer fmt.Println("main done")
	go func() {
		ch <- 42
		fmt.Println("send done")
		defer fmt.Println("goroutine1 done")
		time.Sleep(2 * time.Second)
	}()
	go func() {
		v := <-ch
		fmt.Println("receive done")
		fmt.Println(v)
		defer fmt.Println("goroutine2 done")
		time.Sleep(4 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}
