package leetcode

import "container/heap"

// 定一个非空的整数数组，返回其中出现频率前 k 高的元素。

//

// 示例 1:

// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
// 示例 2:

// 输入: nums = [1], k = 1
// 输出: [1]
//

// 提示：

// 你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
// 你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
// 你可以按任意顺序返回答案。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/top-k-frequent-elements
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type HeapNode struct {
	Val   int
	Times int
}

type TopKHeap []*HeapNode

func (a TopKHeap) Len() int           { return len(a) }
func (a TopKHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TopKHeap) Less(i, j int) bool { return a[i].Times < a[j].Times }
func (a *TopKHeap) Pop() interface{} {
	item := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return item
}

func (a *TopKHeap) Push(item interface{}) {
	*a = append(*a, item.(*HeapNode))
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n] += 1
	}

	size := 0
	kHeap := TopKHeap{}
	for val, times := range countMap {
		node := &HeapNode{
			Val:   val,
			Times: times,
		}
		if size < k {

			heap.Push(&kHeap, node)
			size++
		} else {
			if times > kHeap[0].Times {
				heap.Pop(&kHeap)
				heap.Push(&kHeap, node)
			}
		}
	}

	res := []int{}
	for _, node := range kHeap {
		res = append(res, node.Val)
	}
	return res
}

// type Feq struct {
// 	val int
// 	cnt int
// }
// type FeqMinHeap []Feq

// func (pq *FeqMinHeap) Len() int {
// 	return len(*pq)
// }
// func (pq *FeqMinHeap) Less(i, j int) bool {
// 	return (*pq)[i].cnt < (*pq)[j].cnt
// }
// func (pq *FeqMinHeap) Swap(i, j int) {
// 	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
// }

// func (pq *FeqMinHeap) Push(x interface{}) {
// 	*pq = append(*pq, x.(Feq))
// }

// func (pq *FeqMinHeap) Pop() interface{} {
// 	n := len(*pq) - 1
// 	x := (*pq)[n]
// 	*pq = (*pq)[:n]
// 	return x
// }
// func (pq *FeqMinHeap) Peek() Feq {
// 	return (*pq)[0]
// }

// func topKFrequentV2(nums []int, k int) []int {
// 	cntMap := make(map[int]int)
// 	for _, num := range nums {
// 		cntMap[num]++
// 	}
// 	minHeap := &FeqMinHeap{}
// 	for key, val := range cntMap {
// 		heap.Push(minHeap, Feq{
// 			val: key,
// 			cnt: val,
// 		})
// 		if minHeap.Len() > k {
// 			heap.Pop(minHeap)
// 		}
// 	}
// 	res := make([]int, k)
// 	i := 0
// 	for minHeap.Len() > 0 {
// 		res[i] = minHeap.Pop().(Feq).val
// 		i++
// 	}
// 	return res
// }
