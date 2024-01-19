package leet_code

// 回文数
// https://leetcode.cn/problems/palindrome-number/description/

// 第一种做法 颠倒后面一半的数字
func IsPalindrome(x int) bool {
	// 为负数就不可能是回文数
	// 如果个位数刚好是0 那本身就必须是0才可能是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	// 记录将 x 的右边的一半数字颠倒过来的结果
	temp := 0
	// 每次都 x % 10 取个位数的数字，temp 往前进一位 + 当前获取到的个位数 直到 >= x，则代表已经取到一半了
	for x > temp {
		temp = temp*10 + x%10
		// 去掉已经取出来的个位数
		x = x / 10
	}
	// x的长度是偶数的情况下 两个数相等则为回文数
	// x的长度是奇数的情况下 temp会比 x 多一位数，是中位数， temp / 10去掉该中位数和 x 相等则为回文数
	return x == temp || x == temp/10
}

// 直接完全颠倒整个数字
func IsPalindrome2(x int) bool {
	// 为负数就不可能是回文数
	// 如果个位数刚好是0 那本身就必须是0才可能是回文数
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	// 记录将整个半数字颠倒过来的结果
	x1, temp := x, 0
	for x1 > 0 {
		temp = temp*10 + x1%10
		// 去掉已经取出来的个位数
		x1 = x1 / 10
	}
	return x-temp == 0
}
