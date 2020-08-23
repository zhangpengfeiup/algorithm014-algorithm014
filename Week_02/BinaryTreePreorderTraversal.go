/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }

    result := make([]int, 0)
    stack := make([]*TreeNode, 0)

    for root != nil || len(stack) != 0 {
        for root != nil {
            result = append(result, root.Val)
            stack = append(stack, root)
            root = root.Left
        }
        // pop
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        root = node.Right
    }

    return result
}
