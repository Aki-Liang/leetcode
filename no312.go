package leetcode

// 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

// 现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

// 求所能获得硬币的最大数量。

// 说明:

// 你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
// 0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
// 示例:

// 输入: [3,1,5,8]
// 输出: 167
// 解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//      coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/burst-balloons
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//动态规划法

// 在数组两端添加两个值为1的边界元素方便计算

// 创建二维数组作为dp，存储数组上区间的最大金币数量

// 如dp[i][j]中保存数组上下标i到下标j的开区间内戳气球能获得的最大金币数量

// 假设每个区间最后一个被戳爆的球下标为k，可得出状态转移方程
// dp[i][j] =      dp[i][k]       +      dp[k][j]       + nums[i]*nums[k]*nums[j]
// ------------[i,k]区间最大金币数    [k,j]区间最大金币数   因为前面的已经被戳爆了，所以这里nums[i]和nums[j]就成了k的相邻气球

// nums    	1,  3,  1,  5,  8,  1

// dp  		0,  1,  2,  3,  4,  5
// 		 0  0,  0,  3, 30,159,167
// 		 1  0,  0,  0, 15,135,159
// 		 2  0,  0,  0,  0, 40, 40
// 		 3  0,  0,  0,  0,  0, 40
// 		 4  0,  0,  0,  0,  0,  0
// 		 5  0,  0,  0,  0,  0,  0

func maxCoins(nums []int) int {
	n := len(nums)
	rec := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		rec[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += rec[i][k] + rec[k][j]
				rec[i][j] = max(rec[i][j], sum)
			}
		}
	}
	return rec[0][n+1]
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
