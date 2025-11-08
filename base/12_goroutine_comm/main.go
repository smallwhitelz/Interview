package main

import (
	"fmt"
	"sync"
)

// 协程之间如何通信
// 1. 如果是临界变量要保护，那么就是用sync包下的一些操作
// 2. 如果是goroutine之间发送数据，信号传递，使用channel更合适

func add() {
	var count int
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			lock.Lock()
			count++
			lock.Unlock()
		}
	}()
	wg.Wait()
	fmt.Println(count)
}

func TestChannel(ch chan int) {
	ch <- 1
}

func main() {
	//ch := make(chan int)
	//go TestChannel(ch)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//c := <-ch
	//fmt.Println(c)
	//wg.Done()
	//wg.Wait()
	add()
}
