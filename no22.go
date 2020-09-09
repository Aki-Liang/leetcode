package leetcode

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 示例：

// 输入：n = 3
// 输出：[
//        "((()))",
//        "(()())",
//        "(())()",
//        "()(())",
//        "()()()"
//      ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/generate-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//采用回溯法
func generateParenthesis(n int) []string {
	res := []string{}
	if n == 0 {
		return res
	}
	backTrackParenthesis(n, 1, 0, []byte{'('}, &res)
	return res
}

func backTrackParenthesis(n, p, q int, strByte []byte, res *[]string) {
	if n == q {
		tmp := make([]byte, len(strByte))
		copy(tmp, strByte)
		*res = append(*res, string(tmp))
	}

	if p < n {
		strByte = append(strByte, '(')
		backTrackParenthesis(n, p+1, q, strByte, res)
		strByte = strByte[:len(strByte)-1]
	}
	if q < p {
		strByte = append(strByte, ')')
		backTrackParenthesis(n, p, q+1, strByte, res)
		strByte = strByte[:len(strByte)-1]
	}
}
