package leetcode

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

// 示例 1：

// 输入: "babad"
// 输出: "bab"
// 注意: "aba" 也是一个有效答案。
// 示例 2：

// 输入: "cbbd"
// 输出: "bb"

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-palindromic-substring
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return s
	}

	dp := make([]bool, n)
	maxLen := 0
	left := 0
	right := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				dp[j] = true
			} else if s[i] == s[j] && (j+1 == i || (j+1 < i && dp[j+1])) {

				dp[j] = true
				l := i - j + 1
				if l > maxLen {
					maxLen, left, right = l, j, i
				}
			} else {
				dp[j] = false
			}
		}
	}
	return s[left : right+1]
}
