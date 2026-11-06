import "slices"

func groupAnagrams(strs []string) [][]string {
	hMap := make(map[string][]string)
	var retArr [][]string
	for _, i := range strs{
		s := []rune(i)
		slices.Sort(s)
		hMap[string(s)] = append(hMap[string(s)], i)
	}
	for j := range hMap{
		retArr = append(retArr ,hMap[j])
	}
	return retArr
}
