package leetcode

import "math"

//给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。
//
//你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
//
// 
//
//示例 1：
//
//输入：[1, 2, 2, 3, 1]
//输出：2
//解释：
//输入数组的度是2，因为元素1和2的出现频数最大，均为2.
//连续子数组里面拥有相同度的有如下所示:
//[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
//最短连续子数组[2, 2]的长度为2，所以返回2.
//示例 2：
//
//输入：[1,2,2,3,1,4,2]
//输出：6
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/degree-of-an-array
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func findShortestSubArray(nums []int) int {
    left, right, count := make(map[int]int),make(map[int]int),make(map[int]int)
    maxCount := math.MinInt64
    for index, num := range nums {
        c, ok := count[num]
        if ok {
            right[num] = index

        } else {
            left[num] = index
        }
        c++
        count[num] = c
        if maxCount < c {
            maxCount =c
        }
    }
    res := math.MaxInt64
    for num, c := range count {
        if c == maxCount {
            res = min(res, right[num]-left[num]+1)
        }
    }

    return res
}