import "slices"
func minEatingSpeed(piles []int, h int) int {
	if h == len(piles){
		return slices.Max(piles)
	}

	l := 1
	r := slices.Max(piles)
	
	res := slices.Max(piles)

	for l <= r{
		m := (l + r)/ 2

		count := 0

		for _ , i := range piles{
			count = count + int(math.Ceil(float64(i)/float64(m)))
			fmt.Println(i/m)
		}

		if count > h{
			l = m + 1
		}else{
			r = m - 1
			res = min(res,m)
		}
	}
	return res
}
