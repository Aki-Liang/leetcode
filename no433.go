package leetcode

// 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。

// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。

// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。

// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。

// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。

// 注意:

// 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
// 所有的目标基因序列必须是合法的。
// 假定起始基因序列与目标基因序列是不一样的。
// 示例 1:

// start: "AACCGGTT"
// end:   "AACCGGTA"
// bank: ["AACCGGTA"]

// 返回值: 1
// 示例 2:

// start: "AACCGGTT"
// end:   "AAACGGTA"
// bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]

// 返回值: 2
// 示例 3:

// start: "AAAAACCC"
// end:   "AACCCCCC"
// bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]

// 返回值: 3

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/minimum-genetic-mutation
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var mutationMap = map[uint8][3]string{
	'A': [...]string{"T", "G", "C"},
	'C': [...]string{"T", "G", "A"},
	'T': [...]string{"A", "G", "C"},
	'G': [...]string{"T", "A", "C"},
}

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func minMutation(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	if start == end {
		return 0
	}
	count := 0
	isUsed := make([]bool, len(bank))

	startQueue := []string{start}
	endQueue := []string{end}
	for len(startQueue) > 0 {
		count++
		l := len(startQueue)
		for _, curr := range startQueue {
			for i := 0; i < len(curr); i++ {
				for _, c := range mutationMap[curr[i]] {
					gene := curr[:i] + c + curr[i+1:]
					if idx := idxOf(gene, endQueue); idx != -1 {
						return count
					}
					if idx := idxOf(gene, bank); idx != -1 && !isUsed[idx] {
						isUsed[idx] = true
						startQueue = append(startQueue, gene)
					}
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return -1
}

// func minMutation(start string, end string, bank []string) int {
// 	possible := []string{"A", "C", "G", "T"}
// 	stepMap := make(map[string]int)
// 	for _, w := range bank {
// 		stepMap[w] = math.MaxInt32
// 	}
// 	queue := []string{start}
// 	stepMap[start] = 0i
// 	for len(queue) > 0 {
// 		newQ := []string{}
// 		n := len(queue)
// 		for i := 0; i < n; i++ {
// 			w := queue[i]
// 			if w == end {
// 				return stepMap[w]
// 			}
// 			step, _ := stepMap[w]
// 			for i := 0; i < len(w); i++ {
// 				for _, p := range possible {
// 					tmp := w[:i] + p + w[i+1:]
// 					tmpStep, ok := stepMap[tmp]
// 					if ok && step+1 <= tmpStep {
// 						newQ = append(newQ, tmp)
// 						stepMap[tmp] = step + 1
// 					}
// 				}
// 			}
// 		}
// 		queue = newQ
// 	}
// 	return -1
// }

// func minMutationV2(start string, end string, bank []string) int {
// 	res := -1
// 	wdict := map[string]int{}
// 	for idx, word := range bank {
// 		wdict[word] = idx
// 	}
// 	if _, ok := wdict[end]; !ok {
// 		return res
// 	}
// 	if _, ok := wdict[start]; !ok {
// 		bank = append(bank, start)
// 		wdict[start] = len(bank) - 1
// 	}
// 	n := len(bank)
// 	edge := make([][]int, n)
// 	for i := 0; i < n-1; i++ {
// 		for j := i + 1; j < n; j++ {
// 			if checkIfTrans(bank[i], bank[j]) {
// 				edge[i] = append(edge[i], j)
// 				edge[j] = append(edge[j], i)
// 			}
// 		}
// 	}

// 	cost := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		cost[i] = math.MaxInt32
// 	}
// 	cost[wdict[start]] = 0
// 	queue := [][]int{[]int{wdict[start]}}

// 	for i := 0; i < len(queue); i++ {
// 		now := queue[i]
// 		last := now[len(now)-1]
// 		for _, idx := range edge[last] {
// 			if bank[idx] == end {
// 				return len(now) + 1
// 			}
// 			if cost[last]+1 <= cost[idx] {
// 				queue = append(queue, append(now, idx))
// 				cost[idx] = cost[last] + 1
// 			}
// 		}
// 	}

// 	return res
// }

// func checkIfTrans(w1, w2 string) bool {
// 	diffCount := 0
// 	for i := 0; i < len(w1); i++ {
// 		if w1[i] != w2[i] {
// 			diffCount++
// 			if diffCount > 1 {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
