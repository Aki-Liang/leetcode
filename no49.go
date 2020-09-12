package leetcode

import (
	"sort"
	"strings"
)

// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

// 示例:

// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// 说明：

// 所有输入均为小写字母。
// 不考虑答案输出的顺序。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/group-anagrams
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func groupAnagrams(strs []string) [][]string {
	resMap := make(map[string]int)
	res := make([][]string, 0)
	index := 0
	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sortedStr := strings.Join(chars, "")
		idx, ok := resMap[sortedStr]
		if ok {
			res[idx] = append(res[idx], str)
		} else {
			res = append(res, []string{str})
			resMap[sortedStr] = index
			index++
		}
	}

	return res
}

// func groupAnagrams(strs []string) [][]string {
// 	resMap := make(map[string][]string)
// 	for _, str := range strs {
// 		strBytes := []byte(str)
// 		sort.Slice(strBytes, func(i, j int) bool {
// 			return strBytes[i] < strBytes[j]
// 		})
// 		sortedStr := string(strBytes)
// 		resList, ok := resMap[sortedStr]
// 		if ok {
// 			resList = append(resList, str)
// 		} else {
// 			resList = []string{str}

// 		}
// 		resMap[sortedStr] = resList
// 	}

// 	res := [][]string{}
// 	for _, resList := range resMap {
// 		res = append(res, resList)
// 	}

// 	return res
// }
