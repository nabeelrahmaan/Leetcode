func maxProfit(prices []int) int {
    low := prices[0]
    high := 0
    for _,val := range prices{
        if low>val{
            low=val
        }
        if val-low>high{
            high=val-low
        }
 
    }
    return high
}