package leetcode

import "math"

// 给定一个非空二维矩阵 matrix 和一个整数 k，找到这个矩阵内部不大于 k 的最大矩形和。

// 示例:

// 输入: matrix = [[1,0,1],[0,-2,3]], k = 2
// 输出: 2
// 解释: 矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
// 说明：

// 矩阵内的矩形区域面积必须大于 0。
// 如果行数远大于列数，你将如何解答呢？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows := len(matrix)
	cols := len(matrix[0])
	max := math.MinInt32
	for l := 0; l < cols; l++ { //枚举左边界
		rowSum := make([]int, rows)
		for r := l; r < cols; r++ { // 枚举右边界
			for i := 0; i < rows; i++ { // 按每一行累计到 dp
				rowSum[i] += matrix[i][r]
			}
			max = MaxInt(max, dpmax(rowSum, k))
			if max == k {
				return k
			}
		}
	}
	return max
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dpmax(arr []int, k int) int {
	rollMax, rollSum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if rollSum > 0 {
			rollSum += arr[i]
		} else {
			rollSum = arr[i]
		}
		if rollSum > rollMax {
			rollMax = rollSum
		}
	}
	if rollMax <= k {
		return rollMax
	}

	max := math.MinInt32
	for i := 0; i < len(arr); i++ {
		sum := 0
		for r := i; r < len(arr); r++ {
			sum += arr[r]
			if sum > max && sum <= k {
				max = sum
			}
			if max == k {
				return k
			}
		}
	}

	return max
}

// func maxSumSubmatrix(matrix [][]int, k int) int {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return 0
// 	}
// 	row, col := len(matrix), len(matrix[0])
// 	for r := 0; r < row; r++ {
// 		for c := 1; c < col; c++ {
// 			matrix[r][c] += matrix[r][c-1]
// 		}
// 	}
// 	res := math.MinInt32
// 	for left := 0; left < col; left++ {
// 		for right := left; right < col; right++ {
// 			for top := 0; top < row; top++ {
// 				sum := 0
// 				for bottom := top; bottom < row; bottom++ {
// 					if left == 0 {
// 						sum += matrix[bottom][right]
// 					} else {
// 						sum += matrix[bottom][right] - matrix[bottom][left-1]
// 					}
// 					if sum <= k && sum > res {
// 						res = sum
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return res
// }
