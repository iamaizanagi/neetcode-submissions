func maxProfit(prices []int) int {
    i := 0
    j := 1

    res := 0

    for j < len(prices){
        m := prices[j] - prices[i]
        res = max(res, m)

        if prices[j] < prices[i]{
            i += 1
        }else{
            j += 1
        }
    }
    return res
}
