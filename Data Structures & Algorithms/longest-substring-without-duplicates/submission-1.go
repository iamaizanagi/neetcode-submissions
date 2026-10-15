func lengthOfLongestSubstring(s string) int {
	i , res := 0, 0
	set := make(map[byte]bool)

	for j := 0; j < len(s); j++ {
		for set[s[j]]{
			delete(set, s[i])
			i += 1
		}
		set[s[j]] = true
		res = max(res, (j - i +1 ))
	}
	return res
}
