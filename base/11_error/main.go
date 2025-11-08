package main

import "fmt"

// recover的作用其实是避免我们写大量的if err != nil这种垃圾代码的
// 直接在程序启动入口去捕获所有可能会发生异常的情况，在一个入口统一处理
// 结合方式一般是 panic+recover

func main() {
	// 调用过程中的 panic 被 recover 捕获后程序继续执行
	fmt.Println("Program start")

	// 使用 defer + recover 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// 模拟 panic
	fmt.Println("Before panic")
	panic("Something went wrong!") // 触发 panic

	fmt.Println("After panic") // 这行不会被执行 Unreachable code 遥不可及的代码
}
