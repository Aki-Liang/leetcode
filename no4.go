package leetcode
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//示例 3：
//
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
//示例 4：
//
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
//示例 5：
//
//输入：nums1 = [2], nums2 = []
//输出：2.00000
// 
//
//提示：
//
//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    totalLen := len(nums1)+len(nums2)
    if totalLen%2 == 1{
        midIndex := totalLen/2
        return float64(getKthElem(nums1, nums2, midIndex + 1))
    } else {
        midIndex1,  midIndex2:= totalLen/2-1, totalLen/2
        return float64(getKthElem(nums1, nums2, midIndex1 + 1)+ getKthElem(nums1, nums2, midIndex2 + 1))/2.0
    }
    return 0
}

func getKthElem(nums1 []int, nums2 []int, k int) int {
    idx1, idx2 := 0,0
    for {
        if idx1 == len(nums1) {
            return nums2[idx2+k-1]
        }
        if idx2 == len(nums2) {
            return nums1[idx1+k-1]
        }
        if k == 1 {
            return min(nums1[idx1], nums2[idx2])
        }
        half := k/2
        newIdx1 := min(idx1+half, len(nums1)) - 1
        newIdx2 := min(idx2+half, len(nums2)) - 1
        pivot1, pivot2 := nums1[newIdx1], nums2[newIdx2]
        if pivot1 <= pivot2 {
            k -= newIdx1-idx1+1
            idx1 = newIdx1+1
        } else {
            k -= newIdx2 - idx2 + 1
            idx2 = newIdx2 +1
        }
    }
    return 0
}

//func min(a,b int) int {
//    if a < b {
//        return a
//    }
//    return b
//}
