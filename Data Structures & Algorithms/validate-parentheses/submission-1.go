func isValid(s string) bool {
	stack := []rune{}

	stackMap := map[rune]rune{')':'(', '}':'{', ']':'['}
	//fmt.Println(stackMap)

	for _, ch := range s{
		if open, exists := stackMap[ch]; exists{
			if len(stack) > 0{
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top != open{
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
