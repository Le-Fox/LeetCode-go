package leetcodego

func searchInsert(nums []int, target int) int {
	result := 0
	for i, num := range nums {
		if num >= target {
			result = i
			break
		} else if num < target {
			result = len(nums)
		}
	}

	return result
}
