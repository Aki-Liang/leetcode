package leetcode

// 机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：

// -2：向左转 90 度
// -1：向右转 90 度
// 1 <= x <= 9：向前移动 x 个单位长度
// 在网格上有一些格子被视为障碍物。

// 第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])

// 机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。

// 返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。

//

// 示例 1：

// 输入: commands = [4,-1,3], obstacles = []
// 输出: 25
// 解释: 机器人将会到达 (3, 4)
// 示例 2：

// 输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
// 输出: 65
// 解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/walking-robot-simulation
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type Point struct {
	x int
	y int
}

func robotSim(commands []int, obstacles [][]int) int {
	//阻碍点集合
	obMap := make(map[Point]bool)
	for _, ob := range obstacles {
		p := Point{
			x: ob[0],
			y: ob[1],
		}
		obMap[p] = true
	}
	move := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d, x, y, res := 0, 0, 0, 0
	for _, c := range commands {
		if c == -1 {
			//右转90
			d++
			if d == 4 { //转了一圈
				d = 0
			}
		} else if c == -2 {
			d--
			if d == -1 {
				d = 3
			}
		} else {
			for ; c > 0; c-- {
				p := Point{
					x: x + move[d][0],
					y: y + move[d][1],
				}
				if _, ok := obMap[p]; !ok {
					x += move[d][0]
					y += move[d][1]
				}
			}
		}
		res = max(res, x*x+y*y)
	}

	return res
}
