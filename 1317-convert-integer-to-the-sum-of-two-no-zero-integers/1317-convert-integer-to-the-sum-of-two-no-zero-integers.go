func getNoZeroIntegers(n int) []int {
		for i := 1; i < n; i++ {
			a := i
            b := n-i
            if !strings.Contains(strconv.Itoa(a), "0") && !strings.Contains(strconv.Itoa(b), "0"){
                return []int{a,b}
            }
		}
return nil
}