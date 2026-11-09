func twoSum(numbers []int, target int) []int {
	first, last := 0, len(numbers) - 1
	res := []int{}

	for first < last{
		if (numbers[first] + numbers[last]) == target {
			res = append(res, first+1, last+1)
			break
		}else if (numbers[first] + numbers[last]) < target{
			first += 1
		}else{
			last -= 1
		}
	}
	return res
}
