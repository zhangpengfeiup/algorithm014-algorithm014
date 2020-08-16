/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // 思路：建立一个空节点
 // 循环判断两个链表是否走到结尾，哪一个值小将节点加入到 空节点链表的后面
 // 判断两个链表是否还有节点，将剩余的节点将入到后面

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{0,nil};

    cur := dummy

     for l1 != nil && l2 != nil {
        if (l1.Val < l2.Val) {
            cur.Next = l1
            l1 = l1.Next
            cur = cur.Next
        }else {
            cur.Next = l2
            l2 = l2.Next
            cur = cur.Next
        }
    }

    if l1 != nil {
        cur.Next = l1
    }

    if l2 != nil {
        cur.Next = l2
    }
    
    return dummy.Next
}
