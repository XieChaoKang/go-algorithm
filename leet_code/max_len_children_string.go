package leet_code

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

func LengthOfLongestSubstring2(s string) int {
	stringIndexMap := map[byte]int{}
	right, maxLen := -1, 0
	length := len(s)
	for i := 0; i < length; i++ {
		if i != 0 {
			// 左移一位时 就要删除上一个节点 从窗口队列中移除掉
			delete(stringIndexMap, s[i-1])
		}
		// 右指针开始往下滑动 直到队伍结束或者碰到和开头重复的节点则停止
		for right+1 < length && stringIndexMap[s[right+1]] == 0 {
			stringIndexMap[s[right+1]]++
			right++
		}
		// 左指针到右指针之间的字符就是无重复子串
		if right-i+1 > maxLen {
			maxLen = right - i + 1
		}
	}
	return maxLen
}

// 思路 利用滑动窗口的思想 窗口不停的往右边扩展 每扩展一位就判断当前窗口队列是否已经存在该字符 有则窗口直接缩减到之前旧的位置+1 接着继续该过程 直到结束 中途每次窗口变更都判断一下当前队列大小 永远保留一个max长度值
func LengthOfLongestSubstring(s string) int {
	var childrenString []rune
	maxLen := 0
	for _, i := range s {
		index := indexOf(childrenString, i)
		if index > -1 {
			childrenString = childrenString[index+1:]
		}
		childrenString = append(childrenString, i)
		if len(childrenString) > maxLen {
			maxLen = len(childrenString)
		}
	}
	if len(childrenString) > maxLen {
		maxLen = len(childrenString)
	}
	return maxLen
}

func indexOf(array []rune, target rune) int {
	for index, r := range array {
		if r == target {
			return index
		}
	}
	return -1
}
