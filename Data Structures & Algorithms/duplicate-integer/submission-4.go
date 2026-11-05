func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]int)

	for _, i := range nums{
		hashMap[i]++
	}

	for i := range hashMap{
		if hashMap[i] > 1{
			return true
		}
	}
	return false
}
