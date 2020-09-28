package leetcode

// 给定一个正整数 n，返回长度为 n 的所有可被视为可奖励的出勤记录的数量。 答案可能非常大，你只需返回结果mod 109 + 7的值。

// 学生出勤记录是只包含以下三个字符的字符串：

// 'A' : Absent，缺勤
// 'L' : Late，迟到
// 'P' : Present，到场
// 如果记录不包含多于一个'A'（缺勤）或超过两个连续的'L'（迟到），则该记录被视为可奖励的。

// 示例 1:

// 输入: n = 2
// 输出: 8
// 解释：
// 有8个长度为2的记录将被视为可奖励：
// "PP" , "AP", "PA", "LP", "PL", "AL", "LA", "LL"
// 只有"AA"不会被视为可奖励，因为缺勤次数超过一次。
// 注意：n 的值不会超过100000。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/student-attendance-record-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//动态规划

//每一层在上一层的末尾+A +P +L
// 有题目可知，符合条件的出勤序列有以下六种情况
// A0L0 没有A，末尾没有L
// A0L1 没有A，末尾1个连续L
// A0L2 没有A，末尾2个连续L
// A1L0 1个A，末尾没有L
// A1L1 1个A，末尾1个连续L
// A1L2 1个A，末尾2个连续L

// 每层在以上六种情况的序列末尾+A +P +L

//状态转移方程
// +P，末尾添加P之后状态全部变成末尾没有连续L
// dp[i][A0L0] = dp[i-1][A0L0] +dp[i-1][A0L1] +dp[i-1][A0L2]
// dp[i][A1L0] = dp[i-1][A1L0] +dp[i-1][A1L1] +dp[i-1][A1L2]

// +L，只有末尾没有两个连续L的情况才能追加L
// dp[i][A0L1] = dp[i-1][A0L0]
// dp[i][A0L2] = dp[i-1][A0L1]
// dp[i][A1L1] = dp[i-1][A1L0]
// dp[i][A1L2] = dp[i-1][A1L1]

// +A, 只有没有A的情况才能追加A,末尾追加A之后全部变成尾部没有L的状态
// 因为前面已经计算过+p的情况，所以本层A1L0的情况= 上层A1L*末尾加P + 上层A0L*末尾加A
// dp[i][A1L0] = dp[i-1][A0L0] +dp[i-1][A0L1] +dp[i-1][A0L2] +dp[i][A1L0]
//             = dp[i][A0L0] + dp[i][A1L0]

func checkRecord(n int) int {
	M := 1000000007
	a0l0, a0l1, a0l2, a1l0, a1l1, a1l2 := 1, 0, 0, 0, 0, 0
	for i := 0; i <= n; i++ {
		//+P
		a0l0Tmp := (a0l0 + a0l1 + a0l2) % M
		//+L
		a0l2 = a0l1
		a0l1 = a0l0
		a0l0 = a0l0Tmp
		//+P +A
		a1l0Tmp := (a0l0 + a1l0 + a1l1 + a1l2) % M
		//+L
		a1l2 = a1l1
		a1l1 = a1l0
		a1l0 = a1l0Tmp
	}

	return a1l0
}
