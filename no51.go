package leetcode

import "math/bits"

// n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

// 上图为 8 皇后问题的一种解法。

// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

//

// 示例：

// 输入：4
// 输出：[
//  [".Q..",  // 解法 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // 解法 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// 解释: 4 皇后问题存在两个不同的解法。
//

// 提示：

// 皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-queens
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func solveNQueens(n int) [][]string {
	solutions := [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(queens, n, 0, 0, 0, 0, &solutions)
	return solutions
}

func solve(queens []int, n, row, columns, diagonals1, diagonals2 int, solutions *[][]string) {
	if row == n {
		board := generateBoard(queens, n)
		*solutions = append(*solutions, board)
		return
	}
	availablePositions := ((1 << n) - 1) & (^(columns | diagonals1 | diagonals2))
	for availablePositions != 0 {
		position := availablePositions & (-availablePositions)
		availablePositions = availablePositions & (availablePositions - 1)
		column := bits.OnesCount(uint(position - 1))
		queens[row] = column
		solve(queens, n, row+1, columns|position, (diagonals1|position)>>1, (diagonals2|position)<<1, solutions)
		queens[row] = -1
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

// func solveNQueens(n int) [][]string {
// 	res := [][]string{}
// 	queens := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		queens[i] = -1
// 	}
// 	columns := map[int]bool{}
// 	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
// 	backtrack(queens, n, 0, columns, diagonals1, diagonals2, &res)
// 	return res
// }

// func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool, res *[][]string) {
// 	if row == n {
// 		board := generateBoard(queens, n)
// 		*res = append(*res, board)
// 		return
// 	}

// 	for i := 0; i < n; i++ {
// 		if columns[i] {
// 			continue
// 		}

// 		d1 := row - i
// 		if diagonals1[d1] {
// 			continue
// 		}
// 		d2 := row + i
// 		if diagonals2[d2] {
// 			continue
// 		}
// 		queens[row] = i
// 		columns[i] = true
// 		diagonals1[d1], diagonals2[d2] = true, true
// 		backtrack(queens, n, row+1, columns, diagonals1, diagonals2, res)
// 		queens[row] = -1
// 		delete(columns, i)
// 		delete(diagonals1, d1)
// 		delete(diagonals2, d2)
// 	}
// }

// func generateBoard(queens []int, n int) []string {
// 	board := []string{}
// 	for i := 0; i < n; i++ {
// 		row := make([]byte, n)
// 		for j := 0; j < n; j++ {
// 			row[j] = '.'
// 		}
// 		row[queens[i]] = 'Q'
// 		board = append(board, string(row))
// 	}
// 	return board
// }
