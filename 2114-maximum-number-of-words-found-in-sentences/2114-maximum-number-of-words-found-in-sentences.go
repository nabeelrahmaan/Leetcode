func mostWordsFound(sentences []string) int {
    size := 0
    for _,val := range sentences{
        newarr := strings.Fields(val)
        if len(newarr)>size{
            size=len(newarr)
        }
    }
    return size
}