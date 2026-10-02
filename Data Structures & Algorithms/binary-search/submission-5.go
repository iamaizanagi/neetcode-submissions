func search(nums []int, target int) int {
	l , r := 0 , len(nums) - 1

	if len(nums) == 1 && target == nums[0]{
		return 0
	}
	//m := len(nums) / 2
	for l <= r{
		m := (l + r) / 2
		if nums[m] == target{
			return m
		}else if nums[m] < target{
			l = m + 1
		}else{
			r = m - 1
		}
		m = (l + r) / 2
	}
	return -1 
}
