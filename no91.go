package leetcode

// 一条包含字母 A-Z 的消息通过以下方式进行了编码：

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。

// 示例 1:

// 输入: "12"
// 输出: 2
// 解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
// 示例 2:

// 输入: "226"
// 输出: 3
// 解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/decode-ways
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 动态规划法
// dp中存储译码方法数量，dp[i]存储str[0...i]的方法总数
//s[i]==0, if s[i-1]==1 || s[i-1]==2, 则dp[i]=dp[i-2],否则return 0
// s[i-1]==1, dp[i] = dp[i-1]+dp[i-2]
// s[i-1]==2 && 1<=s[i]<=6, dp[i] = dp[i-1]+dp[i-2]

// index  -1  0 1 2 3 4 5 6
// str        1 2 1 0 1 2 1
// dp      1  1 2 3 2 2 4 6

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, curr := 1, 1 //dp[-1] = dp[0] = 1
	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			curr = curr + pre
		}
		pre = tmp
	}

	return curr
}
