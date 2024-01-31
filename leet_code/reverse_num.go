// https://leetcode.cn/problems/reverse-integer/description/
package leet_code

import "math"

// 思路 直接反转数字 每次都取个位数 不停的往前推一位即可 注意判断反转后的数字的边界
func reverse(x int) int {
	temp := 0
	for x != 0 {
		if temp < math.MinInt32/10 || temp > math.MaxInt32/10 {
			return 0
		}
		temp = temp*10 + x%10
		x /= 10
	}
	return temp
}
