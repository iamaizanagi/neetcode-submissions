
func isPalindrome(s string) bool {
	str := ""
	lowerS := strings.ToLower(s)
	for _ , char := range lowerS{
		if (int(char) >= int('a') && int(char) <= int('z')){
			str = str + string(char)
		} else if (int(char) >= int('0') && int(char) <= int('9') ){
			str = str + string(char)
		}
	}

	fmt.Println(str)

	l := 0
	r := len(str) - 1

	for l < r{
		if str[l] != str[r]{
			return false
		}
		l += 1
		r -= 1
	}
	return true
}
