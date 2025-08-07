package main

import (
	"fmt"
	"runtime"
	"time"
)

// 问题：goroutine什么情况会发⽣内存泄漏？如何避免？
// 严格意义上来说，Go 里面是不存在内存泄露的，毕竟 Go 是垃圾回收语言。
// 因此说 Go 中的内存泄露，其实是指一个对象你预期它很快会被回收，但是最终它过了很久都没有被回收。

//因此最容易引起泄露的就两种情况：
// 切片，尤其是子切片操作；
// goroutine 泄露引发的 goroutine 引用的对象无法被回收的问题；
// 还有无缓冲channel
// mutex并发抢锁也有可能导致

// 切片操作会泄漏
func sliceTest() {
	// 模仿大切片
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	smallS := s[1:4]
	// 这里可以看到新的子切片，底层数组还是8个容量，但是实际上我们用不到了，这样其实就是内存泄漏
	fmt.Printf("val:%v, len:%d, cap:%d", smallS, len(smallS), cap(smallS))

	fmt.Println()
	// 正常玩法：新建一个切片,copy需要的值
	smallS2 := append([]int{}, s[1:4]...)
	fmt.Printf("val:%v, len:%d, cap:%d", smallS2, len(smallS2), cap(smallS2))
}

// leakGoroutine 启动一个永不退出的 goroutine，
// 并在闭包里分配一个很大的切片。
func leakGoroutine() {
	go func() {
		// 在闭包里分配一个 100MB 的大切片
		big := make([]byte, 100<<20) // 100 * 2^20 字节
		// 为了让编译器和 GC 认为 big 正在被“用到”，
		// 我们在死循环里偶尔访问它：
		for i := 0; ; i++ {
			if i%1_000_000 == 0 {
				// 这里只是访问一下，防止被优化掉
				_ = big[0]
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func printMemStats(tag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc = %.2f MB (%.0f KB), HeapObjects = %d\n",
		tag,
		float64(m.Alloc)/1024/1024,
		float64(m.Alloc)/1024,
		m.HeapObjects,
	)
}

func main() {
	//sliceTest()
	// 启动泄露的 goroutine
	leakGoroutine()

	// 等 goroutine 分配完大切片
	time.Sleep(200 * time.Millisecond)

	// 第一次打印：应该看到 ~100MB 已被分配
	printMemStats("Before GC")

	// 强制触发一次垃圾回收
	runtime.GC()

	// 第二次打印：由于那个 goroutine 还在持有 big，所以内存不会被释放
	printMemStats("After GC")

	// 阻塞主 goroutine，不退出程序
	select {}
}
