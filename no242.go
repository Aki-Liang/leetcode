package leetcode

import "sync"

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 示例 1:

// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:

// 输入: s = "rat", t = "car"
// 输出: false
// 说明:
// 你可以假设字符串只包含小写字母。

// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-anagram
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//go 中数组可以直接比较，
func isAnagram(s string, t string) bool {
	//题目中说明只包含小写字母
	bitmapS := [26]int{}
	bitmapT := [26]int{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(s); i++ {
			index := s[i] - 'a'
			bitmapS[index] += 1
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(t); i++ {
			index := t[i] - 'a'
			bitmapT[index] += 1
		}
	}()
	wg.Wait()

	return bitmapS == bitmapT
}

// type bytes []byte

// func (a bytes) Len() int           { return len(a) }
// func (a bytes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a bytes) Less(i, j int) bool { return a[i] < a[j] }

// func isAnagramV2(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	sb := bytes(s)
// 	tb := bytes(t)
// 	sort.Sort(sb)
// 	sort.Sort(tb)

// 	return string(sb) == string(tb)
// }

// func isAnagramV3(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	countMap := make(map[byte]int)
// 	for _, b := range []byte(s) {
// 		count, _ := countMap[b]
// 		countMap[b] = count + 1
// 	}
// 	for _, b := range []byte(t) {
// 		count, _ := countMap[b]
// 		countMap[b] = count - 1
// 	}

// 	for _, value := range countMap {
// 		if value != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	//bitmap 方式，假定只有26个小写字母
// 	sBitmap := [26]int{}
// 	tBitMap := [26]int{}

// 	wg := sync.WaitGroup{}

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < len(s); i++ {
// 			j := s[i] - 'a'
// 			sBitmap[j]++
// 		}
// 	}()

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < len(t); i++ {
// 			j := t[i] - 'a'
// 			tBitMap[j]++
// 		}
// 	}()

// 	wg.Wait()

// 	//go 的数组是可以比较的
// 	return sBitmap == tBitMap
// }
