func maxSlidingWindow(nums []int, k int) []int {
    if len(nums)==0 {
        return []int{}
    }
    var arr []int
    var queue []int
    for i:=0 ; i<len(nums) ; i++ {
        if len(queue)>0 && queue[0]<=i-k {
            queue = queue[1:]
        }
        for len(queue)>0 && nums[queue[len(queue)-1]]<nums[i] {
            queue = queue[:len(queue)-1]
        }
       queue = append(queue, i)
       if i>= k-1 {
        arr = append(arr, nums[queue[0]])
       }
    }
    return arr
}