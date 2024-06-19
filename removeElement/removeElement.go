package removeelement

func removeElement(nums []int, val int) int {
	ret := 0
	for i := 0; i < len(nums); {
		if nums[i] != val {
			nums[ret] = nums[i]
			ret++
		}
		i++
	}
	return ret
}
