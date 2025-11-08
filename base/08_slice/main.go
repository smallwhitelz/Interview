package main

import "fmt"

// 问题：slice 和数组有什么不同？
// 数组定长，默认初始化为0值，函数传参的时候传的是一个副本，修改参数不会改变原有的数组
// slice切片可以扩容，默认是nil，需要显示初始化，函数传参也传了底层数组的指针，
// 所以改变slice形参也会改变底层数组的值，同时在传参的过程中要注意slice的扩容问题
// slice有点类似功能缩减版的ArrayList，缺少随机插入删除的功能，底层是共享数组

func array() {
	n := [3]int{1, 2, 3}
	arr := make([]int, 0, 10) // 如果长度也是10，那么相当于会用int的零值进行初始化10个元素的
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	fmt.Println(n, arr) // [1 2 3] [1 2 3 4]
	copyArray(n, arr)
	fmt.Println(n, arr) // [1 2 3] [1 9 3 4]
}

func copyArray(n [3]int, a []int) {
	n[1] = 5
	fmt.Println(n) // [1 5 3]
	a[1] = 9
	fmt.Println(a) //[1 9 3 4]
}

func remove() {
	n := []int{1, 2, 3, 4, 5, 6, 7}
	// 删除4
	n = append(n[:3], n[4:]...)
	// 即便删除了，底层数组容量还是7
	fmt.Println(n, cap(n))

}

func main() {
	//remove()
	array()
}
