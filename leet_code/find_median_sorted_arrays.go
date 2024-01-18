package leet_code

// 寻找两个正序数组的中位数
// https://leetcode.cn/problems/median-of-two-sorted-arrays/description/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getMiddleElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getMiddleElement(nums1, nums2, midIndex1+1)+getMiddleElement(nums1, nums2, midIndex2+1)) / 2.0
	}
}

// 思路：因为两个数组都是保证有序的 所以只需要用分别遍历两个数组即可
// 两个指针分别记录两个数组遍历到的下标 小的那个向前移动 当其中一个数组已经遍历完了 那就直接移动另外一个数组 补齐剩余需要移动的次数即为目标 或者遍历次数达到了需要移动的次数即可
func getMiddleElement(nums1, nums2 []int, needMoveNum int) int {
	idx1, idx2 := 0, 0
	for {
		// 只需要移动一次的时候直接判断大小即可
		if needMoveNum == 1 {
			return min(nums1[idx1], nums1[idx2])
		}
		// 其中一个数组已经遍历完了 那就直接移动另外一个数组 补齐剩余需要移动的次数即为目标
		if idx1 == len(nums1) {
			return nums2[idx2+needMoveNum-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+needMoveNum-1]
		}
		half := needMoveNum / 2
		// 注意边界判定
		newIdx1 := min(idx1+half, len(nums1)) - 1
		newIdx2 := min(idx2+half, len(nums2)) - 1
		cur1, cur2 := nums1[newIdx1], nums2[newIdx2]
		// 小的那个向前移动 更新剩余需要遍历的次数
		if cur1 < cur2 {
			needMoveNum -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		} else {
			needMoveNum -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
