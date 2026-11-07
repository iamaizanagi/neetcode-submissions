func topKFrequent(nums []int, k int) []int {
	hMap := make(map[int]int)

	for _ ,i := range nums{
		hMap[i]++
	}
	
	var retArr []int

	maxArr := make([][]int, len(nums)+1)

	for j, v := range hMap{
		maxArr[v] = append(maxArr[v], j)
	}
	
	for l := len(maxArr)-1; l > 0 ; l--{
		if len(retArr) == k{
			break
		}
		for m := 0; m < len(maxArr[l]); m++{
			retArr = append(retArr, maxArr[l][m])
		}
	}
	return retArr
}
