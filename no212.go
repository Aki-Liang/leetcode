package leetcode

// 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

// 示例:

// 输入:
// words = ["oath","pea","eat","rain"] and board =
// [
//   ['o','a','a','n'],
//   ['e','t','a','e'],
//   ['i','h','k','r'],
//   ['i','f','l','v']
// ]

// 输出: ["eat","oath"]
// 说明:
// 你可以假设所有输入都由小写字母 a-z 组成。

// 提示:

// 你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
// 如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。

type TrieNode struct {
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	result := []string{}
	root := &TrieNode{}
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); i++ {
			dfsTrie(i, j, board, root, &result)
		}
	}
	return result
}

func dfsTrie(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {
		//访问过的节点或者字典中没有
		return
	}
	node = node.children[c-'a']
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	dfsTrie(i+1, j, board, node, result)
	dfsTrie(i, j+1, board, node, result)
	dfsTrie(i-1, j, board, node, result)
	dfsTrie(i, j-1, board, node, result)
	board[i][j] = c
}
