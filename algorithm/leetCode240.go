package main

import (
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, val := range row {
			if val == target {
				return true
			}
		}
	}
	return false
}

// 二分查找
// 因为都是升序排列，所以二分查找每一列即可
func searchMatrixV1(matrix [][]int, target int) bool {
	for _, row := range matrix {
		idx := sort.SearchInts(row, target)
		if idx < len(row) && row[idx] == target {
			return true
		}
	}
	return false
}

// 45°去看该矩阵，就是一个二叉搜索树
// 从矩阵 matrix 左下角元素（索引设为 (i, j) ）开始遍历，并与目标值对比：
// 当 matrix[i][j] > target 时，执行 i-- ，即消去第 i 行元素。
// 当 matrix[i][j] < target 时，执行 j++ ，即消去第 j 列元素。
// 当 matrix[i][j] = target 时，返回 true ，代表找到目标值。
// 若行索引或列索引越界，则代表矩阵中无目标值，返回 false 。
func searchMatrixV2(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j <= len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
