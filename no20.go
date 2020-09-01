package leetcode

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-parentheses
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isValid(s string) bool {
	dic := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		p, ok := dic[s[i]]
		if ok {
			if len(stack) == 0 || p != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

// func isValid(s string) bool {
// 	if "" == s {
// 		return true
// 	}
// 	if len(s) < 2 {
// 		return false
// 	}
// 	myList := list.New()
// 	dic := map[byte]byte{
// 		')': '(',
// 		'}': '{',
// 		']': '[',
// 	}
// 	str := []byte(s)
// 	for _, b := range str {
// 		pair, ok := dic[b]
// 		if ok {
// 			if myList.Len() == 0 {
// 				return false
// 			}
// 			e := myList.Back()
// 			v := e.Value.(byte)
// 			if pair != v {
// 				return false
// 			}
// 			myList.Remove(e)
// 		} else {
// 			myList.PushBack(b)
// 		}
// 	}

// 	if myList.Len() > 0 {
// 		return false
// 	}
// 	return true
// }
