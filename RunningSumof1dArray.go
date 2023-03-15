package leetcodego

func runningSum(nums []int) []int {
	currentValue := 0
	for i := 0; len(nums) > i; i++ {
		nums[i] = nums[i] + currentValue
		currentValue = nums[i]
	}
	return nums
}
