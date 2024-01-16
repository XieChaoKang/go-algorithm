package leet_code

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
// https://leetcode.cn/problems/two-sum/description/

func TwoSum(nums []int, target int) []int {
	var res []int
	tempMap := make(map[int]int, len(nums)/2*2)
	for index, num := range nums {
		v, ok := tempMap[num]
		if ok {
			res = append(res, v, index)
			break
		}
		tempMap[target-num] = index
	}
	return res
}
