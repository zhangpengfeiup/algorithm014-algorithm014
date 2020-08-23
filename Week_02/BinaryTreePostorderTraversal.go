/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
     // 通过lastVisit 标识右子节点是否已经弹出
     if root == nil {
         return nil
     }

     result := make([]int, 0)
     stack := make([]*TreeNode, 0)
     var lastVisit *TreeNode

     for root != nil || len(stack) != 0 {
         for root != nil {
             stack = append(stack, root)
             root = root.Left
         }
         // 先看看，并不一定弹出
         node := stack[len(stack) - 1]
         if node.Right == nil || node.Right == lastVisit {
             // pop and insert result
             result = append(result, node.Val)
             stack = stack[:len(stack) - 1]

             // update lastVisit
             lastVisit = node
         }else {
             root = node.Right
         }
     }
     
     return result
}
