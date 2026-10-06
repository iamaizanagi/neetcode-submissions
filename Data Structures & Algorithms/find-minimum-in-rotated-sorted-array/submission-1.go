func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	res := nums[0]
	for l <= r{

		if nums[l] < nums[r]{
			res = min(res,nums[l])
			break
		}
		m := (l + r) / 2
		res = min(res,nums[m])
		if nums[m] >= nums[l]{
			res = min(res,nums[m])
			l = m + 1
		}else {
			//fmt.Println("mid > r", nums[mid])
			//min = nums[m]
			
			r = m - 1
		}
		m = (l+r)/2
	}

	return res
}
