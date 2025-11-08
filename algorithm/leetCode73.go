package main

import "fmt"

// 矩阵置零

// 首先遍历该数组一次，如果某个元素为 0，那么就将该元素所在的行和列所对应标记数组的位置置为 true。
// 最后我们再次遍历该数组，用标记数组更新原数组即可。
func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}

func main() {
	nums := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(nums)
	fmt.Println(nums)
}
