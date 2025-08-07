package main

import (
	"context"
	"fmt"
	"time"
)

//问题：select 是如何选择分支的？
// 随机的
// 1. 锁定case中所有的channel
// 2. 按照随机顺序检测，如果case中的channel已经就绪，有读写操作的就执行，否则返回
// 3. 如果所有的case都没有准备好，那么就走default分支
// 4. 如果没有default分支并且case都没准备好，那么就让goroutine阻塞，并且加到所有case的channel中等待被唤醒
// goroutine唤醒后，返回对应的case的index，执行对应的代码
// 所以可以看出来，select的设计和map有点像，不能强依赖于顺序
// 而且“同时”这个从严格意义来说是无法真的同时，因为总会有先后顺序，所以这里同时的意思是当开始执行select的时候，很多的case
// 已经就绪了
// 还有就是当goroutine阻塞的时候，会把这个goroutine加到所有case的channel中等待被唤醒，唤醒后，就要从这些channel的等待队列移除

func Select() {
	// 创建一个带 500ms 超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 查询截止时间和是否存在
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Printf("未到达截止时间，deadline: %v\n", deadline)
	} else {
		fmt.Println("此 Context 没有设置截止时间")
	}

	// 启动一个工作协程，模拟 1s 的耗时操作
	go func() {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("子协程：工作完成")
		case <-ctx.Done():
			// 当超时时，Done 会被关闭，可通过 Err 得知具体原因
			fmt.Printf("子协程：检测到 Done，Err: %v\n", ctx.Err())
		default:
			fmt.Println("默认分支")
		}
	}()

	// 等待子协程完成或超时
	time.Sleep(3 * time.Second)
}

func main() {
	Select()
}
