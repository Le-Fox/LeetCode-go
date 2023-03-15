package leetcodego

func pivotIndex(nums []int) int {
	pivot, n := -1, len(nums)

	for i := 1; n > i; i++ {
		nums[i] += nums[i-1]
	}
	if nums[n-1]-nums[0] == 0 {
		return 0
	}

	for i := 1; n > i; i++ {
		left := nums[i-1]
		right := nums[n-1] - nums[i]

		if left == right {
			pivot = i
			break
		}
	}

	return pivot

}
