package leetcode

// n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

// 上图为 8 皇后问题的一种解法。

// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。

// 示例:

// 输入: 4
// 输出: 2
// 解释: 4 皇后问题存在如下两个不同的解法。
// [
//  [".Q..",  // 解法 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // 解法 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/n-queens-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func totalNQueens(n int) int {
	count := 0
	solveII(n, 0, 0, 0, 0, &count)
	return count
}

func solveII(n, row, columns, diagonals1, diagonals2 int, count *int) {
	if row == n {
		*count++
		return
	}

	bits := (^(columns | diagonals1 | diagonals2)) & ((1 << n) - 1)
	for bits != 0 {
		p := bits & (-bits)      //取到最低位的1
		bits = bits & (bits - 1) //在p的位置上放上皇后
		solveII(n, row+1, columns|p, (diagonals1|p)<<1, (diagonals2|p)>>1, count)
	}

}
