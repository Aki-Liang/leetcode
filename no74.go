package leetcode

// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 示例 1:

// 输入:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 3
// 输出: true
// 示例 2:

// 输入:
// matrix = [
//   [1,   3,  5,  7],
//   [10, 11, 16, 20],
//   [23, 30, 34, 50]
// ]
// target = 13
// 输出: false

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/search-a-2d-matrix
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// func searchMatrix(matrix [][]int, target int) bool {
// 	nums := []int{}
// 	for _, items := range matrix {
// 		nums = append(nums, items...)
// 	}
// 	return searchList(nums, target)
// }

// func searchList(nums []int, target int) bool {
// 	if len(nums) == 0 {
// 		return false
// 	}
// 	index := len(nums) / 2
// 	if target == nums[index] {
// 		return true
// 	} else if target < nums[index] {
// 		return searchList(nums[:index], target)
// 	} else {
// 		return searchList(nums[index+1:], target)
// 	}
// }

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if target > matrix[m-1][n-1] || target < matrix[0][0] {
		return false
	}
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)>>1
		r := mid / n
		c := mid % n
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return false
}
