package main

import (
	"context"
	"fmt"
	"time"
)

// 问题：介绍⼀下 Go 中的 context 有什么用？
// 用作数据传递
// 用作超时传递
// 用作信号传递（超时本身可以看作是一种信号）

// ctx类似于别的语言中的ThreadLocal，提供线程安全的读写操作
// 但是ctx有一些额外的问题，就是用Value获取数据的时候返回值是一个any，我们每次需要断言获取到的值的类型
// ctx本身是不可变的对象，每次创建都是一个新的副本，所以上游无法感知到下游对ctx进行了什么改变

func foo() {
	ctx := context.WithValue(context.Background(), "traceId", "11111")
	bar(ctx)
	fmt.Println(ctx.Value("traceId")) // 但是这里输出还是111111
}
func bar(ctx context.Context) {
	val := context.WithValue(ctx, "traceId", "22222").Value("traceId")
	fmt.Println(val) // 这里已经改为了222222
}

func timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("任务完成")
	case <-ctx.Done():
		// 超时或被取消时 ctx.Err() 会返回具体原因
		fmt.Printf("超时或取消: %v\n", ctx.Err())
	}
}

func signal() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Printf("收到信号，原因: %v\n", ctx.Err())
	}(ctx)
	// 模拟一些工作
	time.Sleep(500 * time.Millisecond)
	// 发送取消信号
	cancel()
	// 等待子协程打印
	time.Sleep(100 * time.Millisecond)
}

func main() {
	foo()
	//timeout()
	//signal()
}
