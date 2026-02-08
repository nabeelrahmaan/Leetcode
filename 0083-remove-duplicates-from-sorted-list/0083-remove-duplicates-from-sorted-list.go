/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    data := head
    for data.Next!=nil{
        if data.Val == data.Next.Val{
            data.Next = data.Next.Next
        }else{
            data = data.Next
        }

    }
    return head
}