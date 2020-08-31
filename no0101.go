package leetcode

// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

// 示例 1：

// 输入: s = "leetcode"
// 输出: false
// 示例 2：

// 输入: s = "abc"
// 输出: true
// 限制：

// 0 <= len(s) <= 100
// 如果你不使用额外的数据结构，会很加分。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/is-unique-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func isUnique(astr string) bool {

	var high uint64
	var low uint64

	for i := 0; i < len(astr); i++ {
		ch := astr[i]
		if ch >= 64 {
			//高64位
			bit := uint64(1 << (ch - 64))
			if 0 != (high & bit) {
				return false
			}
			high |= bit
		} else {
			//低64位
			bit := uint64(1 << ch)
			if 0 != (low & bit) {
				return false
			}
			low |= bit
		}
	}
	return true
}
