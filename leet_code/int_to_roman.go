package leet_code

// https://leetcode.cn/problems/integer-to-roman/description/

func intToRoman(num int) string {
	var romanConf = []struct {
		value       int
		romanSymbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var res []byte
	for _, conf := range romanConf {
		for num >= conf.value {
			num -= conf.value
			res = append(res, conf.romanSymbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(res)
}
