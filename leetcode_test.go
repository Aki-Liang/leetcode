package leetcode

import (
	"fmt"
	"testing"
)

func Test_tribonacci(t *testing.T) {
	res := tribonacci(4)
	if 4 != res {
		panic(res)
	}
}

func Test_findUnsortedSubarray(t *testing.T) {
	res := findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
	//res := findUnsortedSubarray([]int{1, 2, 3, 4})
	if 5 != res {
		panic(res)
	}
}

func Test_plusOne(t *testing.T) {
	res := plusOne([]int{9, 9, 9, 9})
	fmt.Println(res)
}
