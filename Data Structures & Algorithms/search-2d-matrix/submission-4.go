func searchMatrix(matrix [][]int, target int) bool {
	rows, col := len(matrix), len(matrix[0])

	top, bot := 0, rows - 1

	for top <= bot{
		mid := (top + bot) / 2

		if target > matrix[mid][(len(matrix[bot])-1)]{
			top = mid + 1
		}else if target < matrix[mid][0]{
			bot = mid - 1
		}else{
			break
		}
	}

	if !(top <= bot){
		return false
	}

	mid := (top + bot) / 2

	l, r := 0, col - 1

	for l <= r {
		m := (l + r ) / 2

		if target > matrix[mid][m]{
			l = m + 1
		}else if target < matrix[mid][m]{
			r = m - 1
		}else{
			return true
		}
	}
	return false
}
