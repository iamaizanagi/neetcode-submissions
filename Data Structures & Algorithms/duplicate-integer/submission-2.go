func hasDuplicate(nums []int) bool {
    hMap := make(map[int]int)

	for _ , value := range nums{
		hMap[value]++

		if hMap[value] > 1{
			return true
		}
	}
	return false
}
