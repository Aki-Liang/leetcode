package leetcode

import "math"

//给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
//
//上图子矩阵左上角 (row1, col1) = (2, 1) ，右下角(row2, col2) = (4, 3)，该子矩形内元素的总和为 8。
//
// 
//
//示例：
//
//给定 matrix = [
//[3, 0, 1, 4, 2],
//[5, 6, 3, 2, 1],
//[1, 2, 0, 1, 5],
//[4, 1, 0, 1, 7],
//[1, 0, 3, 0, 5]
//]
//
//sumRegion(2, 1, 4, 3) -> 8
//sumRegion(1, 1, 2, 2) -> 11
//sumRegion(1, 2, 2, 4) -> 12
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/range-sum-query-2d-immutable
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type NumMatrix struct {
    matrix [][]int
    sumMatrix [][]int
}


func NumMatrixConstructor(matrix [][]int) NumMatrix {
    res := NumMatrix{
        matrix: matrix,
    }
    n := len(matrix)
    if n > 0 {
        sumMatrix := make([][]int , n)
        for i:=0; i<n; i++ {
            m := len(matrix[i])
            sumMatrix[i] = make([]int, m)
            sumMatrix[i][0]= matrix[i][0]
            for j:=1; j<m; j++ {
                sumMatrix[i][j] = matrix[i][j] + sumMatrix[i][j-1]
            }
        }
        res.sumMatrix = sumMatrix
    }

    return res
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    n, m := len(this.matrix), len(this.matrix[0])
    if row1 < 0 || row1 > n || row2 < 0 || row2 > n || col1 < 0 || col1 > m || col2 <0 || col2 > m {
        return math.MinInt64
    }

    sum := 0
    for i:= row1; i<= row2; i++ {
        if col1 > 0 {
            sum += this.sumMatrix[i][col2] - this.sumMatrix[i][col1-1]
        } else {
            sum += this.sumMatrix[i][col2]
        }
    }
    return sum
}
