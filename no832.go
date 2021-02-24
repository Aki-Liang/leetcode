package leetcode

//给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。
//
//水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。
//
//反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。

//示例 1：
//
//输入：[[1,1,0],[1,0,1],[0,0,0]]
//输出：[[1,0,0],[0,1,0],[1,1,1]]
//解释：首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
//然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
//示例 2：
//
//输入：[[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
//输出：[[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
//解释：首先翻转每一行: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]]；
//然后反转图片: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
// 
//
//提示：
//
//1 <= A.length = A[0].length <= 20
//0 <= A[i][j] <= 1
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/flipping-an-image
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func flipAndInvertImage2(A [][]int) [][]int {
    res := [][]int{}
    for _, line := range A {
        l := []int{}
        for i:= len(line)-1; i>=0; i-- {
            if line[i] == 0 {
                l = append(l, 1)
            } else {
                l = append(l, 0)
            }
        }
        res = append(res, l)
    }
    return res
}

func flipAndInvertImage(A [][]int) [][]int {
    for i, line := range A {
        m := len(line)
        for j := 0; j<(m+1)/2; j++ {
            A[i][j], A[i][m-1-j] = A[i][m-1-j]^1, A[i][j]^1
        }
    }
    return A
}
