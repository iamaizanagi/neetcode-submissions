func isValid(s string) bool {
    stack := []rune{}
	stackMap := map[rune]rune{ '}':'{', ')':'(', ']':'[' }

	for _, ch := range s {
		if val, exists := stackMap[ch]; exists{
			if len(stack) > 0{
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != val{
					return false
				}
			}else{
				return false
			}
		}else{
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
