func twoSum(nums []int, target int) []int {
    hMap := make(map[int]int)
	var retArr []int

	for i , value := range nums{
		hMap[value] = i
	}

	for j := 0; j < len(nums); j++{
		rem := target - nums[j]

		if hMap[rem] != 0 && hMap[rem] != j {
			//fmt.Println(rem)
			retArr = append(retArr, j)
			retArr = append(retArr, hMap[rem])
			break
		}
	}
	return retArr
}
