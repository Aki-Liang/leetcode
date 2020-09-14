package leetcode

// 实现 pow(x, n) ，即计算 x 的 n 次幂函数。

// 示例 1:

// 输入: 2.00000, 10
// 输出: 1024.00000
// 示例 2:

// 输入: 2.10000, 3
// 输出: 9.26100
// 示例 3:

// 输入: 2.00000, -2
// 输出: 0.25000
// 解释: 2-2 = 1/22 = 1/4 = 0.25
// 说明:

// -100.0 < x < 100.0
// n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/powx-n
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//快速幂，每次把上一次的结果平方
//如果要计算 x^64,可以按照
//x -> x^2 -> x^4 -> x^8 -> x^{16} -> x^{32} -> x^{64}

//快速幂 + 迭代
// x^{77} = x∗x^4∗x^8∗x^64
// 1,4,8,64正好对应了 7777 的二进制表示 (1001101)中的每个 1

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	ans := 1.0
	x_contribute := x
	for n > 0 {
		if n%2 == 1 {
			//当前位为1
			ans *= x_contribute
		}
		x_contribute *= x_contribute
		//舍弃最低位
		n = n / 2

	}
	return ans
}
