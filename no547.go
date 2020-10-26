package leetcode

// 班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

// 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

// 示例 1：

// 输入：
// [[1,1,0],
//  [1,1,0],
//  [0,0,1]]
// 输出：2
// 解释：已知学生 0 和学生 1 互为朋友，他们在一个朋友圈。
// 第2个学生自己在一个朋友圈。所以返回 2 。
// 示例 2：

// 输入：
// [[1,1,0],
//  [1,1,1],
//  [0,1,1]]
// 输出：1
// 解释：已知学生 0 和学生 1 互为朋友，学生 1 和学生 2 互为朋友，所以学生 0 和学生 2 也是朋友，所以他们三个在一个朋友圈，返回 1 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/friend-circles
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return 0
	}
	parent := make([]int, n) //对称矩阵
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] == 1 && i != j {
				union(parent, i, j)
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			count++
		}
	}
	return count
}

func find(parent []int, i int) int {
	root := i
	for parent[root] != root {
		root = parent[root]
	}
	for parent[i] != i { //路径压缩
		i, parent[i] = parent[i], root
	}
	return i
}

func union(parent []int, x, y int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	if xSet != ySet {
		parent[xSet] = ySet
	}
}
