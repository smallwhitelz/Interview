package main

import "fmt"

func add(m map[int]int) {
	m[1] = 3
	fmt.Println("add 方法内", m)
}

func main() {
	m := make(map[int]int)
	m[1] = 2
	fmt.Println(m)
	add(m)
	fmt.Println(m)
}
