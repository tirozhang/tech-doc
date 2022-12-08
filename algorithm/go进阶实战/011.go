package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("channel is closed")
				return
			}
			fmt.Println(a)
		}
	}()

	close(ch)
	fmt.Println("main goroutine exit")
	time.Sleep(time.Second * 100)
}
