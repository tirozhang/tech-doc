package main

import (
	"fmt"
	"sync"
)

// 交替打印数字和字母
func main() {
	// 1. 创建两个channel
	// 2. 创建两个goroutine
	// 3. 通过select语句实现交替打印

	number, letter := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	wait.Add(1)
	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}
		}
	}()
	number <- true
	wait.Wait()

}
