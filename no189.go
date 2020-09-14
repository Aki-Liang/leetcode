package leetcode

import "fmt"

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

// 示例 1:

// 输入: [1,2,3,4,5,6,7] 和 k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右旋转 1 步: [7,1,2,3,4,5,6]
// 向右旋转 2 步: [6,7,1,2,3,4,5]
// 向右旋转 3 步: [5,6,7,1,2,3,4]
// 示例 2:

// 输入: [-1,-100,3,99] 和 k = 2
// 输出: [3,99,-1,-100]
// 解释:
// 向右旋转 1 步: [99,-1,-100,3]
// 向右旋转 2 步: [3,99,-1,-100]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/rotate-array
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	count := 0

	for i := 0; count < l; i++ {
		prev := nums[i]
		curr := i
		for {
			curr = (curr + k) % l
			prev, nums[curr] = nums[curr], prev
			count++
			if curr == i {
				break
			}
		}
	}

	fmt.Println(nums)
}

// func rotate(nums []int, k int) {
// 	l := len(nums)
// 	k = k % l

// 	count := 0 //移动到位需要每个元素移动一次，所以这里记录元素移动次数
// 	// prev := 0  //记录前一个被占位的元素
// 	for i := 0; count < l; i++ {
// 		curr := nums[i] //当前要挪走的元素
// 		index := i
// 		for {
// 			index = (index + k) % len(nums) //目的位置的下标
// 			curr, nums[index] = nums[index], curr
// 			count++
// 			if index == i { //表示已经全部交换完毕
// 				break
// 			}
// 		}
// 	}
// 	fmt.Println(nums)
// }

// func rotate(nums []int, k int) {
// 	l := len(nums)
// 	k = k % l
// 	count := 0
// 	for i := 0; count < l; i++ {
// 		idx := i
// 		next := nums[idx]
// 		for {
// 			idx = (idx + k) % len(nums)
// 			nums[idx], next = next, nums[idx]
// 			count++
// 			if idx == i {
// 				break
// 			}
// 		}
// 	}
// 	fmt.Println(nums)
// }

// func rotate(nums []int, k int) {
// 	length := len(nums)
// 	k = k % length
// 	count := 0
// 	for i := 0; count < length; i++ {
// 		prev := nums[i]
// 		current := i
// 		for {
// 			next := (current + k) % length
// 			prev, nums[next] = nums[next], prev
// 			count++
// 			current = next
// 			if current == i {
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(nums)
// }
