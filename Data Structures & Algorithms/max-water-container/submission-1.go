func maxArea(heights []int) int {
	l , r := 0 , len(heights) - 1

	res := 0 

	for l < r{
		area := (r - l) * min(heights[l], heights[r])

		res = max(res, area)

		if heights[l] < heights[r]{
			l += 1
		}else{
			r -= 1
		}
	}
	return res
}
