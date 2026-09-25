func isAnagram(s string, t string) bool {
	hsMap := make(map[byte]int)
	htMap := make(map[byte]int)

	if len(s) != len(t){
		return false
	}
	
	for key := range s{
		hsMap[s[key]]++ 
	}

	for tvalue := range t{
		htMap[t[tvalue]]++
	}

	for i := range hsMap{
		if hsMap[i] != htMap[i]{
			return false
		}
	}
	return true
}
