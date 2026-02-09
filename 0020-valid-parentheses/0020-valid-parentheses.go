func isValid(s string) bool {
    temp := map[string]string{
        ")" : "(",
        "}" : "{",
        "]" : "[",
    }
    var stack []string
    for _, val := range s {
        ch := string(val)
        if open, ok := temp[ch]; ok{
            if len(stack)==0{
                return false
            }
            if stack[len(stack)-1] != open {
                return false
            }
            stack = stack[:len(stack)-1]
        }else{
            stack = append(stack, ch)
        }
    }
    return len(stack)==0 
}