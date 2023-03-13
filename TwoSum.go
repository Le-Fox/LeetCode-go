package leetcodego

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		var start = nums[i]
		for j := i + 1; j < len(nums); j++ {
			var end = nums[j]
			if (start + end) == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
