func characterReplacement(s string, k int) int {
    l := 0 
    
    hashMap := make(map[byte]int)

    res := 0
    maxf := 0

    for r := 0; r < len(s); r++{
        hashMap[s[r]]++

        if hashMap[s[r]] > maxf{
            maxf = hashMap[s[r]]
        }

        for (r - l + 1) - maxf > k{
            hashMap[s[l]]--
            l++
        }

        if (r - l + 1) > res{
            res = r - l + 1
        }
    }
    return res
}
