package algorithms

func InsertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j := i
		tmp := nums[j]
		for j > 0 && nums[j-1] > tmp {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = tmp
	}
}
