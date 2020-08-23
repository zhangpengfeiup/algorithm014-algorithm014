/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    stack := make([]*TreeNode, 0)

    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 取出节点并pop
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        result = append(result, node.Val)
        root = node.Right
    }

    return result
}
