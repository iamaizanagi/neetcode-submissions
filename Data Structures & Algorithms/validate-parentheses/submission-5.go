func isValid(s string) bool {
    var stack []rune
	strMap := map[rune]rune{
		')':'(',
		'}':'{',
		']':'[',
	}
	if len(s) == 1{
		return false
	}

	for _, i := range s{
		if val, exists := strMap[i]; exists{
			if len(stack) > 0{
				top := stack[len(stack)-1]
				if top != val{
					return false
				}
			}else{
				return false
			}
			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, i)
		}
	}
	return len(stack) == 0
}
