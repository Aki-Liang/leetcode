package leetcode

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。

// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

// 请找出其中最小的元素。

// 你可以假设数组中不存在重复元素。

// 示例 1:

// 输入: [3,4,5,1,2]
// 输出: 1
// 示例 2:

// 输入: [4,5,6,7,0,1,2]
// 输出: 0

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] <= nums[right] {
			//证明这个区间上是递增的
			return nums[left]
		}
		//这个区间被旋转过，这里获取中点
		mid := left + (right-left)>>1
		if nums[left] <= nums[mid] {
			//在右边找
			left = mid + 1
		} else {
			//在左边找
			right = mid
		}
	}
	return -1
}
