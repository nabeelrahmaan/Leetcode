func isAnagram(s string, t string) bool {
   if len(s) != len(t) {
    return false
   }
   countS, countT := make(map[rune]int), make(map[rune]int)

   for ind, ch := range s {
    countS[ch]++
    countT[rune(t[ind])]++
   }
   for ind, val := range countS {
    if countT[ind] != val {
        return false
    }
   }
   return true  
}