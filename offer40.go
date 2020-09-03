package leetcode

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

// 示例 1：

// 输入：arr = [3,2,1], k = 2
// 输出：[1,2] 或者 [2,1]
// 示例 2：

// 输入：arr = [0,1,2,1], k = 1
// 输出：[0]
//

// 限制：

// 0 <= k <= arr.length <= 10000
// 0 <= arr[i] <= 10000

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//快排
func getLeastNumbers(arr []int, k int) []int {
	return nil
}

// 小顶堆排序法
// type KHeap []int

// func (a KHeap) Len() int           { return len(a) }
// func (a KHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a KHeap) Less(i, j int) bool { return a[i] > a[j] }
// func (a *KHeap) Push(x interface{}) { // add x as element Len()
// 	*a = append(*a, x.(int))
// }
// func (a *KHeap) Pop() interface{} { // remove and return element Len() - 1.
// 	var ret interface{}
// 	l := len(*a)
// 	if l > 0 {
// 		ret = (*a)[l-1]
// 		*a = (*a)[:l-1]
// 	}
// 	return ret
// }

// func getLeastNumbers(arr []int, k int) []int {
// 	h := &KHeap{}
// 	for _, n := range arr {
// 		if len(*h) < k {
// 			heap.Push(h, n)
// 		} else {
// 			heap.Push(h, n)
// 			heap.Pop(h)
// 		}
// 	}

// 	return *h
// }
