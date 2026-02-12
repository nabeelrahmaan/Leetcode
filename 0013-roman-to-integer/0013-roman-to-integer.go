func romanToInt(s string) int {
    mymap := make(map[rune]int)
    mymap['I'] = 1
    mymap['V'] = 5
    mymap['X'] = 10
    mymap['L'] = 50
    mymap['C'] = 100
    mymap['D'] = 500
    mymap['M'] = 1000 

    runes := []rune(s)
    var out int
    for i := 0 ; i< len(runes) ; i++{
        if i != len(runes)-1{
            if mymap[runes[i]] < mymap[runes[i+1]] {
                out += mymap[runes[i+1]] - mymap[runes[i]]
                i++
                continue
            }
        }
      out += mymap[runes[i]]
    }
    return out 
}
 