package leetcode

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

// 你可以认为每种硬币的数量是无限的。

// 示例 1：

// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
// 示例 2：

// 输入：coins = [2], amount = 3
// 输出：-1
// 示例 3：

// 输入：coins = [1], amount = 0
// 输出：0
// 示例 4：

// 输入：coins = [1], amount = 1
// 输出：1
// 示例 5：

// 输入：coins = [1], amount = 2
// 输出：2
//

// 提示：

// 1 <= coins.length <= 12
// 1 <= coins[i] <= 231 - 1
// 0 <= amount <= 231 - 1

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/coin-change
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 动态规划

// coin {1,2,5}
// amount 11
// dp中存储指定amount的最少组成个数
// dp  amount  0,1,2,3,4,5,6,7,8,9,10,11
//     count   0,1,1,2,2,1,2,2,3,3, 2, 3

// 在有1,2,5这三种硬币的情况下，凑成amount的最后一层有三种可能
// dp[amount-1]加面值为1的硬币，dp[amount-2]加面值为2的硬币， dp[amount-5]加面值为5的硬币
// 取其中最小的，个数+1，即可得出amount的最少组成个数
// 推出状态转移方程
// dp[amount] = min(dp[amount-1]+1, dp[amount-2]+1, dp[amount-5]+1)

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1 //初始化成无法达到的值
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}

		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
