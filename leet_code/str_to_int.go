package leet_code

// 将字符串转换成一个 32 位有符号整数
// https://leetcode.cn/problems/string-to-integer-atoi/description/

import (
	"math"
	"strings"
)

func MyAtoi(s string) int {
	IntMax := int(math.Pow(2, 31) - 1)
	IntMin := -int(math.Pow(2, 31))

	// 去掉空格
	s = strings.TrimLeft(s, " ")

	// 判断正负号
	sign := 1
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		sign = -1
		if s[0] == '-' {
			s = s[1:]
		} else {
			s = s[1:]
		}
	}

	num := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		// 判断非数字的字符直接返回
		if ch > 9 {
			break
		}
		// 计算 前移位数 + 当前的数字
		num = num*10 + int(ch)
		// 判断是否超过范围 超过则直接返回对应的范围界限值即可
		if sign*num > IntMax {
			return IntMax
		} else if sign*num < IntMin {
			return IntMin
		}
	}
	return num * sign
}
