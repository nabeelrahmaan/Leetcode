func sumZero(n int) []int {
    var arr[]int
    if n%2!=0{
        arr = append(arr,0)
    }
    for i := 1 ; i<=n/2 ; i++{
        arr = append(arr, i, -i)
    }
    return arr
}