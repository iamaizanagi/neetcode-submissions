func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	min := nums[0]

	for l < r{
		m := (l+r) / 2

		if nums[m] < min{
			min = nums[m]
		}

		if nums[m] < nums[r]{
			 r = m
		}else{
			l = m + 1
		}
	}
	return nums[l]
}
