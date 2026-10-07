func maxProfit(prices []int) int {
    l := 0
    r := 1

    price := 0

    for r < len(prices){

        if prices[l] > prices[r]{
            l += 1
            //r += 1
        }else if prices[r] >= prices[l]{
            if price < (prices[r] - prices[l]){
                price = prices[r] - prices[l]
            }
            r += 1
        }
    }
    return price
}
