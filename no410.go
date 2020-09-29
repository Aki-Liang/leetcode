package leetcode

import "math"

// 给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。

// 注意:
// 数组长度 n 满足以下条件:

// 1 ≤ n ≤ 1000
// 1 ≤ m ≤ min(50, n)
// 示例:

// 输入:
// nums = [7,2,5,10,8]
// m = 2

// 输出:
// 18

// 解释:
// 一共有四种方法将nums分割为2个子数组。
// 其中最好的方式是将其分为[7,2,5] 和 [10,8]，
// 因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/split-array-largest-sum
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 动态规划

// 创建二维数组dp
// dp中存储最大连续子数组和的最小值
// dp[i][j]中保存长度为i的数组分为j段之后能得到的最大连续子数组和的最小值
// 设i的前x个数组成了j-1个连续数组，则max(dp[x][j-1], x到i的子数组和）为i长度数组分为j段的最大连续子数组和
// 通过枚举x，取结果的最小值，即可得出i长度数组分为j段后能得到的最大连续子数组和的最小值

func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	sub := make([]int, n+1) //用来快速计算子数组和
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < min(i, m); j++ {
			for x := 0; x < i; x++ {
				dp[i][j] = min(dp[i][j], max(dp[x][j-1], sub[i]-sub[x]))
			}
		}
	}

	return dp[n][m]
}
