package threesum

import "sort"

func threeSum(nums []int) [][]int {
	lenNums := len(nums)
	ret := make([][]int, 0, 0)
	if lenNums < 3 {
		return ret
	}
	// 排序
	sort.Ints(nums)

	for i := 0; i < lenNums; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} // 去重
		l, r := i+1, lenNums-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				} // 左边去重
				for r < r && nums[r] == nums[r-1] {
					r--
				} // 后边去重
				l++
				r--
			}
			if sum > 0 {
				r--
			}
			if sum < 0 {
				l++
			}
		}
	}
	return ret
}
