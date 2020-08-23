package leetcode

func rangeBitwiseAnd(m int, n int) int {
    for m < n {
        n = n&(n-1)
    }
    return n
}


func rangeBitwiseAndV2(m int, n int) int {
	if m < n {
		return rangeBitwiseAndV2(m/2, n/2) << 1
	} 
	return m
}