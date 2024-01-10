package leet_code

// 删除所有 "AB" 和 "CD" 子串，返回可获得的最终字符串的 最小 可能长度。
// https://leetcode.cn/problems/minimum-string-length-after-removing-substrings/description

// 思路：不断的压进切片中 长度每到 2 或者 2以上就检测最近压进切片的两个字符串是否需要进行删除 这样就只需要遍历一遍即可
func MinLength(s string) int {
	var res []byte
	for i := range s {
		res = append(res, s[i])
		length := len(res)
		if length >= 2 && (string(res[length-2:]) == "AB" || string(res[length-2:]) == "CD") {
			res = res[:length-2]
		}
	}
	return len(res)
}
