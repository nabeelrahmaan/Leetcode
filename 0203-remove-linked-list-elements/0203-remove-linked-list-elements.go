/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }
    if head == nil{
        return nil
    }
    curr := head
    for curr.Next != nil {
        if curr.Next.Val == val {
            curr.Next = curr.Next.Next
            continue
        }
        curr = curr.Next
    }
    return head
}