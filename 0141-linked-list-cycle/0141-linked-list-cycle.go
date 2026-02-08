/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    check := make(map[*ListNode]int)
    if head==nil || head.Next==nil{
        return false
    }
    temp := head
    for temp!=nil{
        if _, ok := check[temp]; ok{
            return true
        }
        check[temp]=temp.Val
        temp = temp.Next
    }
    return false
}