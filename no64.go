package leetcode

// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

// 说明：每次只能向下或者向右移动一步。

// 示例:

// 输入:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// 输出: 7
// 解释: 因为路径 1→3→1→1→1 的总和最小。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-path-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 原数组                  dp
// 1,3,4,8      1,     3+1=4,  4+4=8, 8+8=16
// 3,2,2,4      3+1=4, 4+2=6,  6+2=8, 8+4=12
// 5,7,1,9      4+5=9, 6+7=13, 8+1=9, 9+9=18
// 2,3,2,3      9+2=11,11+3=14,9+2=11,11+3=14

// 状态转移方程
// m>0 && n==0  dp[m][0] = dp[m-1][0]+grid[m][0]
// m==0 && n>0  dp[0][n] = dp[0][n-1]+grid[0][n]
// m>0 && n>0   dp[m][n] = min(dp[m-1][n], dp[m][n-1]) + grid[m][n]

//动态规划解法
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, columns := len(grid), len(grid[0])
	dp := make([]int, columns)
	for m := 0; m < rows; m++ {
		for n := 0; n < columns; n++ {
			if n-1 >= 0 && m-1 >= 0 {
				dp[n] = min(dp[n-1], dp[n]) + grid[m][n]
			} else if n-1 >= 0 && m == 0 {
				dp[n] = dp[n-1] + grid[m][n]
			} else if n == 0 && m-1 >= 0 {
				dp[n] = dp[n] + grid[m][n]
			} else {
				dp[n] = grid[m][n]
			}
		}
	}

	return dp[columns-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func minPathSum(grid [][]int) int {
// 	if len(grid) == 0 || len(grid[0]) == 0 {
// 		return 0
// 	}

// 	rows, columns := len(grid), len(grid[0])
// 	dp := make([][]int, rows)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, columns)
// 	}

// 	dp[0][0] = grid[0][0]
// 	for i := 1; i < rows; i++ {
// 		dp[i][0] = dp[i-1][0] + grid[i][0]
// 	}
// 	for j := 1; j < columns; j++ {
// 		dp[0][j] = dp[0][j-1] + grid[0][j]
// 	}

// 	for i := 1; i < rows; i++ {
// 		for j := 1; j < columns; j++ {
// 			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
// 		}
// 	}

// 	return dp[rows-1][columns-1]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
