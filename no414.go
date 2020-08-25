package leetcode

import "strconv"

// 写一个程序，输出从 1 到 n 数字的字符串表示。

// 1. 如果 n 是3的倍数，输出“Fizz”；

// 2. 如果 n 是5的倍数，输出“Buzz”；

// 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。

// 示例：

// n = 15,

// 返回:
// [
//     "1",
//     "2",
//     "Fizz",
//     "4",
//     "Buzz",
//     "Fizz",
//     "7",
//     "8",
//     "Fizz",
//     "Buzz",
//     "11",
//     "Fizz",
//     "13",
//     "14",
//     "FizzBuzz"
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/fizz-buzz
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		res = append(res, strconv.Itoa(i))
	}
	for i := 2; i < n; i += 3 {
		res[i] = "Fizz"
	}
	for i := 4; i < n; i += 5 {
		res[i] = "Buzz"
	}
	for i := 14; i < n; i += 15 {
		res[i] = "FizzBuzz"
	}
	return res
}

// func fizzBuzz(n int) []string {
// 	res := []string{}
// 	m, n := 1, 1
// 	for i := 1; i <= n; i++ {
// 		m++
// 		n++
// 		if m == 3 && n == 5 {
// 			res = append(res, "FizzBuzz")
// 			m = 0
// 			n = 0
// 		} else if m == 3 {
// 			res = append(res, "Fizz")
// 			m = 0
// 		} else if n == 5 {
// 			res = append(res, "Buzz")
// 			n = 0
// 		} else {
// 			res = append(res, strconv.Itoa(i))
// 		}
// 	}

// 	return res
// }

// func fizzBuzz(n int) []string {
// 	res := []string{}
// 	for i := 1; i <= n; i++ {
// 		v := ""
// 		if 0 == i%3 {
// 			v += "Fizz"
// 		}
// 		if 0 == i%5 {
// 			v += "Buzz"
// 		}
// 		if "" == v {
// 			v = strconv.Itoa(i)
// 		}

// 		res = append(res, v)
// 	}

// 	return res
// }
