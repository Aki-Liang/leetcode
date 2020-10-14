package leetcode

// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

// 此外，你可以假设该网格的四条边均被水包围。

// 示例 1:

// 输入:
// [
// ['1','1','1','1','0'],
// ['1','1','0','1','0'],
// ['1','1','0','0','0'],
// ['0','0','0','0','0']
// ]
// 输出: 1
// 示例 2:

// 输入:
// [
// ['1','1','0','0','0'],
// ['1','1','0','0','0'],
// ['0','0','1','0','0'],
// ['0','0','0','1','1']
// ]
// 输出: 3
// 解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/number-of-islands
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func find(parents []int, i int) int {
	root := i
	for parents[root] != root {
		root = parents[root]
	}

	for parents[i] != i {
		i, parents[i] = parents[i], root
	}

	return i
}

func union(parents []int, i, j int) {
	ri := find(parents, i)
	rj := find(parents, j)
	if ri != rj {
		parents[ri] = rj
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	parents := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		parents[i] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				continue
			}

			pos := i*m + j
			if j+1 < m && grid[i][j+1] == '1' {
				//如果右边是陆地，合并
				union(parents, pos, pos+1)
			}
			if i+1 < n && grid[i+1][j] == '1' {
				//如果下边是陆地，合并
				union(parents, pos, pos+m)

			}
		}
	}

	islands := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && parents[i*m+j] == i*m+j {
				islands++
			}
		}
	}

	return islands
}

// func numIslands(grid [][]byte) int {
// 	count := 0
// 	for line := range grid {
// 		for column := range grid[line] {
// 			if '0' == grid[line][column] { //找到陆地
// 				continue
// 			}
// 			count++
// 			dfs(grid, line, column)
// 		}
// 	}

// 	return count
// }

// func dfs(grid [][]byte, line, column int) {
// 	//找到相连的陆地，然后搞沉没
// 	if line < 0 || line >= len(grid) || column < 0 || column >= len(grid[line]) {
// 		return
// 	}
// 	if grid[line][column] != '1' {
// 		//没有陆地了
// 		return
// 	}
// 	grid[line][column] = '0'
// 	dfs(grid, line+1, column)
// 	dfs(grid, line-1, column)
// 	dfs(grid, line, column+1)
// 	dfs(grid, line, column-1)
// }
