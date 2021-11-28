package sum

// Ints returns the sum of a list of integers
func Ints(nums ...int) int {
	return ints(nums)
}

func ints(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return ints(nums[1:]) + nums[0]
}
