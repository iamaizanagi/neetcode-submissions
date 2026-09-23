import "maps"

func checkInclusion(s1 string, s2 string) bool {
	i , j , res := 0, len(s1)-1, false

	s1map := make(map[byte]int)

	for i := range s1{
		s1map[s1[i]]++
	}
	for j < len(s2){
		winMp := make(map[byte]int)

		for k := i; k <= j ; k++{
			winMp[s2[k]]++
		}
		//fmt.Println(winMp)
		if maps.Equal(s1map,winMp){
			res = true
		}
		i += 1
		j += 1
	}
	return res
}
