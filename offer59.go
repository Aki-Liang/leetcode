package leetcode

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

// 示例:

// 输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
// 输出: [3,3,5,5,6,7]
// 解释:

//   滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
//

// 提示：

// 你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

// 注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//使用双端队列的解法
//在双端队列中，要始终保证队头是队列中最大的值
//在添加一个值之前，比他小的都要被移除掉，然后再添加这个值

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 || k == 0 {
		return res
	}
	//填充滑动窗口
	queue := []int{}
	for i := 0; i < k; i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res = append(res, queue[0])

	for i := k; i < len(nums); i++ {
		if nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		res = append(res, queue[0])
	}

	return res
}

// func maxSlidingWindow(nums []int, k int) []int {
// 	res := []int{}
// 	if len(nums) == 0 || k == 0 {
// 		return res
// 	}
// 	dqueue := []int{}

// 	//填充滑动窗口
// 	for i := 0; i < k; i++ {
// 		for len(dqueue) != 0 && dqueue[len(dqueue)-1] < nums[i] {
// 			dqueue = dqueue[:len(dqueue)-1]
// 		}
// 		dqueue = append(dqueue, nums[i])
// 	}
// 	res = append(res, dqueue[0])
// 	for i := k; i < len(nums); i++ {
// 		if dqueue[0] == nums[i-k] {
// 			dqueue = dqueue[1 : len(dqueue)-1]
// 		}
// 		for len(dqueue) != 0 && dqueue[len(dqueue)-1] < nums[i] {
// 			dqueue = dqueue[:len(dqueue)-1]
// 		}
// 		dqueue = append(dqueue, nums[i])
// 		res = append(res, dqueue[0])
// 	}

// 	return res
// }
