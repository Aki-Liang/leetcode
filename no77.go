package leetcode

// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

// 示例:

// 输入: n = 4, k = 2
// 输出:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/combinations
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func combine(n int, k int) [][]int {
	res := [][]int{}
	if n == 0 || k == 0 || n < k {
		return res
	}
	trackBackCombine(n, 1, k, []int{}, &res) //从1开始
	return res
}

func trackBackCombine(n, m, k int, nums []int, res *[][]int) {
	//先写达成条件
	if len(nums) == k {
		//一个组合完成
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
	}

	for ; m <= n; m++ {
		nums = append(nums, m)
		trackBackCombine(n, m+1, k, nums, res) //拼下一个
		nums = nums[:len(nums)-1]              //还原
	}

}

// func combine(n int, k int) [][]int {
// 	res := [][]int{}
// 	if n <= 0 || k <= 0 || n < k {
// 		return res
// 	}
// 	nums := []int{}
// 	trackBackCombine(n, 1, k, nums, &res)
// 	return res
// }

// func trackBackCombine(n, m, k int, nums []int, res *[][]int) {
// 	if len(nums) == k {
// 		tmp := make([]int, k)
// 		copy(tmp, nums)
// 		*res = append(*res, tmp)
// 		return
// 	}

// 	for ; m <= n; m++ {
// 		nums = append(nums, m)
// 		trackBackCombine(n, m+1, k, nums, res)
// 		nums = nums[:len(nums)-1]
// 	}
// }
