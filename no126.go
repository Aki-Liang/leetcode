package leetcode

import "math"

// 给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：

// 每次转换只能改变一个字母。
// 转换后得到的单词必须是字典中的单词。
// 说明:

// 如果不存在这样的转换序列，返回一个空列表。
// 所有单词具有相同的长度。
// 所有单词只由小写字母组成。
// 字典中不存在重复的单词。
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
// 示例 1:

// 输入:
// beginWord = "hit",
// endWord = "cog",
// wordList = ["hot","dot","dog","lot","log","cog"]

// 输出:
// [
//   ["hit","hot","dot","dog","cog"],
//   ["hit","hot","lot","log","cog"]
// ]
// 示例 2:

// 输入:
// beginWord = "hit"
// endWord = "cog"
// wordList = ["hot","dot","dog","lot","log"]

// 输出: []

// 解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/word-ladder-ii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	idxMap := map[string]int{}
	for i, word := range wordList {
		//建立单词和index的映射关系
		idxMap[word] = i
	}

	if _, ok := idxMap[beginWord]; !ok {
		//将beginword添加到列表尾部，方便寻找
		wordList = append(wordList, beginWord)
		idxMap[beginWord] = len(wordList) - 1
	}
	if _, ok := idxMap[endWord]; !ok {
		return [][]string{}
	}
	n := len(wordList)
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	res := [][]string{}
	queue := [][]int{[]int{idxMap[beginWord]}}
	cost := make([]int, n) //用于标识从begin word变化到指定word时的变化次数
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[idxMap[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		now := queue[i]              //当前层？
		last := now[len(now)-1]      //获取最后一个单词
		if last == idxMap[endWord] { //如果最后一个但是是结束词就填充返回结果集
			tmp := []string{}
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			res = append(res, tmp)
		} else { //不然就还得加
			for _, to := range edges[last] { //遍历最后一个单词能转换的所有单词
				if cost[last]+1 <= cost[to] { //因为初始值是MaxInt32，如果一个单词没有被用过，这里肯定是to的值大于last的值
					//如果cost[to]的值小于cost[last]+1，证明有其他更短的路径可以到cost[to]这个词，就不用再继续转到cost[to]路径
					//如果cost[to]的值大于cost[last]+1，证明当前cost[to]路径是更短的路径，更新之
					cost[to] = cost[last] + 1 //标记为已使用
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)      //新词追加到末尾
					queue = append(queue, tmp) //放到下一轮迭代
				}
			}
		}
	}
	return res
}

func transformCheck(from, to string) bool {
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			return from[i+1:] == to[i+1:]
		}
	}
	return false
}
