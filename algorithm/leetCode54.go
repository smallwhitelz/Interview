package main

// 螺旋矩阵
// 根据题目示例 matrix = [[1,2,3],[4,5,6],[7,8,9]] 的对应输出 [1,2,3,6,9,8,7,4,5]
// 可以发现，顺时针打印矩阵的顺序是 “从左向右、从上向下、从右向左、从下向上” 循环。
// 因此，考虑设定矩阵的 “左、上、右、下” 四个边界，模拟以上矩阵遍历顺序。

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var (
		l = 0                  // 左
		r = len(matrix[0]) - 1 // 右
		t = 0                  // 上
		b = len(matrix) - 1    // 下
		x = 0
	)
	ans := make([]int, len(matrix)*len(matrix[0]))
	for {
		// 从左向右【上边】
		for i := l; i <= r; i++ {
			ans[x] = matrix[t][i]
			x++
		}
		// 上边这一层已经遍历，所以要向内收
		t++
		if t > b {
			break
		}
		// 从上向下【右边】
		for i := t; i <= b; i++ {
			ans[x] = matrix[i][r]
			x++
		}
		r--
		if r < l {
			break
		}
		// 从右向左【下边】
		for i := r; i >= l; i-- {
			ans[x] = matrix[b][i]
			x++
		}
		b--
		if b < t {
			break
		}
		// 从下向上【左边】
		for i := b; i >= t; i-- {
			ans[x] = matrix[i][l]
			x++
		}
		l++
		if l > r {
			break
		}
	}
	return ans
}

func main() {

}
