func trap(height []int) int {
	maxLeft := make([]int, len(height))
	maxRight := make([]int, len(height))
	sum := 0
	mleft := 0
	mright := 0

	for i := 0; i < len(height); i++{
		if i == 0{
			maxLeft[i] = 0
		}else if height[i-1] > mleft{
			mleft = height[i-1]
			maxLeft[i] = mleft
		}else{
			maxLeft[i] = mleft
		}
	}
	//fmt.Println(maxLeft)
	
	for j := (len(height)- 1); j >= 0; j--{
		if j == (len(height)- 1){
			maxRight[j] = 0
		}else if height[j+1] > mright{
			mright = height[j+1]
			maxRight[j] = mright
		}else{
			maxRight[j] = mright
		}
	}
	//fmt.Println(maxRight)
	
	for k := 0; k < len(height); k++{
		water := min(maxLeft[k], maxRight[k]) - height[k]

		if water > 0{
			sum = sum + water
		}
	}
	return sum
}
