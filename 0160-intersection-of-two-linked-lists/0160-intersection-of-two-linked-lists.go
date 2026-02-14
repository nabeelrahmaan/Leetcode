/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
   curr := headA
   temp := headB
    mymap := make(map[*ListNode]int)
   for curr != nil {
        mymap[curr] = curr.Val
        curr = curr.Next
   }

   for temp != nil {
    if _, ok := mymap[temp] ; ok {
        return temp
    }
    temp= temp.Next
   }

return temp
}