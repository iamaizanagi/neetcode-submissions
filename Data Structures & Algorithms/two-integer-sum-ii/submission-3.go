func twoSum(numbers []int, target int) []int {
	res := []int{}

	i := 0
	j := len(numbers) - 1

	for i < j {
		if numbers[i] + numbers[j] > target{
			j -= 1
		}else if numbers[i] + numbers[j] < target{
			i += 1
		}else if numbers[i] + numbers[j] == target{
			//fmt.Println(numbers[i], numbers[j])
			res = append(res, i+1,j+1)
			//res = append(res, numbers[j])
			return res
		}
	}
	return res
}
