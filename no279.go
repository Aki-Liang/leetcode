package leetcode

import "math"

// 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

// 示例 1:

// 输入: n = 12
// 输出: 3
// 解释: 12 = 4 + 4 + 4.
// 示例 2:

// 输入: n = 13
// 输出: 2
// 解释: 13 = 4 + 9.

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/perfect-squares
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//数学法
func numSquaresV2(n int) int {

	// 1770 年，Joseph Louis Lagrange证明了一个定理，称为四平方和定理，也称为 Bachet 猜想，它指出每个自然数都可以表示为四个整数平方和：
	// 	p=a0^2+a1^2+a2^2+a3^2
	// 在 1797 年，Adrien Marie Legendre用他的三平方定理完成了四平方定理，证明了正整数可以表示为三个平方和的一个特殊条件：
	// 	n != 4^k *(8m+7) <=> n =a0^2+a1^2+a2^2

	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}

	if isSquare(n) {
		return 1
	}

	for i := 1; i*i < n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3
}

func isSquare(n int) bool {
	i := int(math.Sqrt(float64(n)))
	return n == i*i
}

//dp 法

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	//dp[i]:构成数i的最小的完全平方和
	//dp[i] = min(dp[i-j*j] + 1 , min);     dp[i] = min(dp[i-k] + 1 ) k:1~ i-k*k >=0
	for i := 1; i <= n; i++ {
		m := math.MaxInt32
		for j := 1; i-j*j >= 0; j++ {
			m = min(dp[i-j*j]+1, m)
		}
		dp[i] = m
	}
	return dp[n]
}
