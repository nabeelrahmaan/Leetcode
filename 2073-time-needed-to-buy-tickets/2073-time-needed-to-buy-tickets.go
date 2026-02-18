func timeRequiredToBuy(tickets []int, k int) int {
    time := 0
    for i := 0 ; tickets[k] > 0 ; i++ {
        if tickets[i] > 0 {
            tickets[i]--
            time++
        }
        if i == len(tickets)-1 {
            i = -1
        }
    }
    return time
}