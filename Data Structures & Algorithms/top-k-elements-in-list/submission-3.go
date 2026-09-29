
func topKFrequent(nums []int, k int) []int {
	hMap := make(map[int]int)
    hArr := make([][]int, len(nums)+1)

	for i := 0; i < len(nums); i++{
		hMap[nums[i]]++
	}

    for key, cnt := range hMap{
        hArr[cnt] = append(hArr[cnt], key)
    }
	res := []int{}
	for j := (len(hArr) - 1); j > 0; j --{
		for _, n := range hArr[j]{
			res = append(res, n)
			if len(res) == k{
				return res
			}
		}
	}
	return res
}
