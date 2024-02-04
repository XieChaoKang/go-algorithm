package leet_code

// https://leetcode.cn/problems/longest-palindromic-substring/description/

// 思路：回文子串在取其中间往两边分别切割出来的字符串刚好是颠倒过来的 也就是以中间为中心点切割开来 分别向左向右遍历过程中 每一个字符都应该是一样的
// 注意：奇数和偶数长度的回文子串的中间是不同的 中位数的概念
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	start, end := 0, 0
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 以当前位置为中心起点 判断是否存在回文子串 也就是右边的剩余字符串和左边的字符串构成回文子串
		// 奇数长度的回文字串的中心就是当前遍历到的位置
		left1, right1 := isPalindrome(s, i, i)
		// 偶数长度的回文字串的中心应该是两个字符串 也就是遍历到的位置 和 遍历到的位置+1
		left2, right2 := isPalindrome(s, i, i+1)
		// 因为是求最长的回文子串 所以计算两个坐标之间的差 如果差 <= 0 也就是意味着当前位置作为中心点时不存在回文子串
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

// 判断是否有一个回文字串
func isPalindrome(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) {
		// 如果当前左右两边分别遍历到的字符不一样就意味着回文字串结束了
		if s[left] != s[right] {
			break
		}
		// 指针分别移动
		left -= 1
		right += 1
	}
	// 返回当前有效的回文字串的两边的指针的下标
	return left + 1, right - 1
}
