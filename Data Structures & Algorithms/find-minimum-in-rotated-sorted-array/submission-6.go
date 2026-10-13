func findMin(nums []int) int {
	l , r := 0 ,len(nums) - 1

	res := nums[l]

	for l <= r {
		m := (l + r)/2

		res = min(res, nums[m])

		if nums[m] >= nums[l]{
			if nums[l] > nums[r]{
				l = m + 1
			}else{
				r = m - 1
			}
		}else{
			r = m - 1
		}
	}
	return res
}
