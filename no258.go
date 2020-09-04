package leetcode

// 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

// 示例:

// 输入: 38
// 输出: 2
// 解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。
// 进阶:
// 你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/add-digits
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 38 = 3*10 + 8
// 第一轮计算3+8，所以先从10个3中拿出一个
// 38 = 3*9+(3+8) = 3*9+11
// 第二轮计算1+1，所以先从10个1中拿出一个
// 38= 3*9+1*10+1 = 3*9 + 1*9 +(1+1) =  3*9 + 1*9 +(2)
// 接下来只要去掉3*9和1*9即可拿到最终结果，模9即可
// 考虑到最终结果为9的情况直接模9会导致结果为0，所以先-1模9，再+1

func addDigits(num int) int {
	return (num-1)%9 + 1
}
