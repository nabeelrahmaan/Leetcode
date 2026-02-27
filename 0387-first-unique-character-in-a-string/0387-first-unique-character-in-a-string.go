func firstUniqChar(s string) int {
    myrune := []rune(s)
    mymap := map[rune]int{}
    for _, val := range myrune {
        mymap[val]++
    }
    for i, val := range myrune {
        if mymap[val] == 1 {
            return i
        } 
    }
    return -1
}