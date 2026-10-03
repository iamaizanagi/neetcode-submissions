func searchMatrix(matrix [][]int, target int) bool {
	for i := 0 ; i < len(matrix); i++{
		l := 0
		r := len(matrix[i]) - 1

		for l <= r{
			m := (l + r )/ 2

			if matrix[i][m] == target{
				return true
			}else if matrix[i][m] > target{
				//fmt.Println("r : ",matrix[i][m])
				r = m - 1
			}else{
				//fmt.Println("l : ", matrix[i][m])
				l = m + 1
			}
		}

	}
	return false
}
