package algorithms

func BubbleSort(nums []int) {
	swaps := -1

	for swaps != 0 {
		swaps = 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swaps++
			}
		}
	}
}
