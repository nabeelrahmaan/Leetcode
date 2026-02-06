func maxProfit(prices []int) int {
    x := 0
    for i := 1 ; i<len(prices) ; i++ {
        if prices[i]>prices[i-1]{
            x += prices[i]-prices[i-1]
        }
    }
    return x
}