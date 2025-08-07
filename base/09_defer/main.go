package main

import "fmt"

// 问题：defer 底层是如何实现的？

// Defer 所有的defer由链表连接起来，添加的时候添加到头部，所以执行的时候也是先从头部开始
// 所以会发现先定义的后执行，后定义的先执行
// defer在Go里面有三种实现方式：开放编码、栈分配和堆分配，Go会依次尝试这三种方式
// 开放编码：把原来在运行时构造以及注册defer对象的过程，直接展开为普通函数的掉哟个，要求：
// 启用了内联优化：简单的函数会直接将其复制到调用处，省去了真实调用函数的开销
// defer语句少于8个
// defer * return语句少于15个

// 栈分配，没有内存逃逸，则分配到栈上
// 堆分配，最后兜底
func Defer() {
	defer fmt.Println(111)
	defer fmt.Println(222)
}

func main() {
	Defer()
}
