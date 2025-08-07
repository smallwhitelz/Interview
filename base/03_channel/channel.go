package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentQueue struct {
	noFull *sync.Cond
}

func init() {
}

// 示例 1：无缓冲 channel
func unbufferedExample() {
	ch := make(chan struct{}) // 无缓冲

	go func() {
		fmt.Println("goroutine：准备发送信号（无缓冲）…")
		ch <- struct{}{} // 阻塞，直到有人接收
		fmt.Println("goroutine：信号已发送（无缓冲）")
	}()

	// 主 goroutine 先睡 1 秒，再接收
	time.Sleep(time.Second)
	fmt.Println("main：准备接收信号（无缓冲）")
	<-ch
	fmt.Println("main：已接收信号（无缓冲）")
}

// 示例 2：缓冲区容量为 1 的 channel
func bufferedExample() {
	ch := make(chan struct{}, 1) // 缓冲 1

	go func() {
		fmt.Println("goroutine：准备发送信号（缓冲 1）…")
		ch <- struct{}{} // 第一次发送不会阻塞
		fmt.Println("goroutine：信号已发送（缓冲 1），即使 main 还没准备好")
	}()

	// 主 goroutine 先睡 1 秒，再接收
	time.Sleep(time.Second)
	fmt.Println("main：准备接收信号（缓冲 1）")
	<-ch
	fmt.Println("main：已接收信号（缓冲 1）")
}

// 问题：channel的发送和接收操作有哪些基本特性?
// 对channel的读写都是线程安全的，不用担心多个goroutine读写同一个channel
// 未初始化的channel读写都会阻塞
func channelExample() {
	// 这是声明一个channel，并不是初始化
	//var ch chan int
	// 这是初始化一个缓冲为1的channel
	ch := make(chan int, 1)
	ch <- 1
	val := <-ch
	fmt.Println(val)
}

// 读取已关闭的channel会先读到有缓冲的数据，读完了后会读到0值，写已关闭的channel会直接panic
// 关闭已经关闭的channel也会直接panic
// 再一个要小心goroutine泄漏
func closeChannel() {
	ch := make(chan int, 1)
	ch <- 1
	val := <-ch
	fmt.Println("未关闭：", val)
	close(ch)
	newVal := <-ch
	fmt.Println("已关闭：", newVal)
	//ch <- 2
	close(ch)
}

func main() {
	//fmt.Println("=== 无缓冲 channel 示例 ===")
	//unbufferedExample()
	//fmt.Println()
	//
	//fmt.Println("=== 缓冲区为 1 的 channel 示例 ===")
	//bufferedExample()
	//channelExample()
	closeChannel()
}
