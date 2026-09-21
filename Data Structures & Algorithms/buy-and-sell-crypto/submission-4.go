func maxProfit(prices []int) int {
    if len(prices) == 1{
        return 0
    }
    i := 0
    j := 1
    profit := 0

    for j < len(prices){
        if prices[i] > prices[j]{
            j += 1
            i = j - 1
        }else{
            //fmt.Println("Entering Second if")
            if profit < (prices[j] - prices[i]){
                profit = prices[j] - prices[i]
            }
            //fmt.Println(profit)
            j += 1
        }
    }
    return profit
}
