package leetcode

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

func trap(height []int) int {
	res := 0
	left := 0
	leftmax := 0
	right := len(height) - 1
	rightmax := 0
	for left <= right {
		if leftmax < rightmax {
			res += getMax(0, leftmax-height[left])
			leftmax = getMax(leftmax, height[left])
			left++
		} else {
			res += getMax(0, rightmax-height[right])
			rightmax = getMax(rightmax, height[right])
			right--
		}
	}

	return res
}

// func trap(height []int) int {
// 	length := len(height)
// 	left := 0
// 	leftMax := 0
// 	right := length - 1
// 	rightMax := 0
// 	var res int
// 	for left <= right {
// 		if leftMax < rightMax {
// 			res += getMax(0, leftMax-height[left])
// 			leftMax = getMax(leftMax, height[left])
// 			left++
// 		} else {
// 			res += getMax(0, rightMax-height[right])
// 			rightMax = getMax(rightMax, height[right])
// 			right--
// 		}
// 	}
// 	return res
// }
