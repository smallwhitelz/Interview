package main

import "C"
import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// 问题：Go什么时候会阻塞？

// Channel 阻塞
func channel() {
	ch := make(chan int) // 无缓冲的channel
	go func() {
		// 这里尝试往 ch 发送数据，但因主 Goroutine 还未接收，会阻塞
		fmt.Println("子 Goroutine: 正在发送 42 到 ch")
		ch <- 42
		fmt.Println("子 Goroutine: 发送完成")
	}()
	// 模拟一点延迟
	// 直接加上这一行，会导致无法输出 fmt.Println("子 Goroutine: 发送完成")
	// 因为主goroutine等待1秒收到数据后立刻就结束了
	// 加了 Sleep 后，主 Goroutine 接收并打印完毕到程序退出的时序，恰好“抢”了子 Goroutine 最后那次打印的调度机会，导致你没看到它。
	time.Sleep(time.Second) // 如果加上这行，主 Goroutine 延迟接收，则子 Goroutine 阻塞时间更长

	// 主 Goroutine 接收数据，唤醒被阻塞的子 Goroutine
	v := <-ch
	fmt.Println("主 Goroutine: 接收到", v)
	// 加上这个就可以看到或者加waitGroup
	//time.Sleep(time.Second)
}

// addMutex 获取锁阻塞
func addMutex() {
	var mu sync.Mutex
	fmt.Println("主goroutine加锁")
	mu.Lock()
	go func() {
		fmt.Println("子 Goroutine: 尝试 mu.Lock()")
		mu.Lock() // 这里会阻塞，直到主 Goroutine 调用 Unlock()
		fmt.Println("子 Goroutine: 锁获取成功")
		mu.Unlock()
	}()

	// 模拟工作
	time.Sleep(time.Second)

	fmt.Println("主 Goroutine: 正在 mu.Unlock()")
	mu.Unlock()
	time.Sleep(time.Second)

	//主goroutine加锁
	//子 Goroutine: 尝试 mu.Lock()
	//主 Goroutine: 正在 mu.Unlock()
	//子 Goroutine: 锁获取成功
}

// io阻塞
func io() {
	fmt.Println("主 Goroutine: 请输入一些文字，然后回车")
	var buf [100]byte
	n, err := os.Stdin.Read(buf[:]) // 读标准输入，直到用户按回车或 EOF
	if err != nil {
		panic(err)
	}
	fmt.Printf("读取到 %d 字节: %q\n", n, buf[:n])
}

// 系统调用也会引发阻塞
// C.sleep(3) 是一个纯粹的系统调用，会阻塞所在的 OS 线程（M）。
// Go 调度器知道 M 阻塞后会脱离 P，将 P 分配给其它 M，继续执行其它 Goroutine。
func system() {
	fmt.Println("调用 C.sleep(3)——当前 OS 线程将被阻塞 3 秒")
	C.sleep(3)
	fmt.Println("C.sleep 返回，继续执行")
}

// busyLoop
// 当一个 Goroutine 连续运行过长（默认大约 10 ms），Go 运行时的抢占机制会强制中断它，检查是否需要做垃圾回收、调度其它 Goroutine，达到了“隐式阻塞”的效果。
// 虽然 busyLoop 中没有任何 IO 或同步操作，运行时也会在合适时机抢占，让出 P 给其它 Goroutine。
func busyLoop() {
	for i := 0; i < 1e9; i++ {
		// 长时间计算，不会主动让出 CPU
	}
}

// 每次调用 Gosched() 后，当前 Goroutine 会被挂起，并放回到可运行队列末尾。
func worker(id int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d: step %d\n", id, i)
		// 主动让出，让调度器去运行别的 Goroutine
		runtime.Gosched()
	}
}

func main() {
	//channel()
	//addMutex()
	io()
}
