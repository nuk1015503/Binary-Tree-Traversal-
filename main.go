package main

import (
	"fmt"
)

const input = "  -"

func main() {
	tmpNode := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6}}
	preStr := preOrder(tmpNode)
	inStr := inOrder(tmpNode)
	fmt.Println(preStr, inStr)
}

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorderTraversal preorderTraversal
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left != nil || root.Right != nil {
		tmp := []int{}
		tmp = append(tmp, root.Val)
		tmp = append(tmp, preorderTraversal(root.Left)...)
		tmp = append(tmp, preorderTraversal(root.Right)...)

		return tmp
	}
	return []int{root.Val}
}

// inorderTraversal inorderTraversal
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left != nil || root.Right != nil {

		tmp := []int{}
		tmp = append(tmp, inorderTraversal(root.Left)...)
		tmp = append(tmp, root.Val)
		tmp = append(tmp, inorderTraversal(root.Right)...)
		return tmp
	}
	return []int{root.Val}
}

// postorderTraversal postorderTraversal
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left != nil || root.Right != nil {

		tmp := []int{}
		tmp = append(tmp, postorderTraversal(root.Left)...)
		tmp = append(tmp, postorderTraversal(root.Right)...)
		tmp = append(tmp, root.Val)
		return tmp
	}
	return []int{root.Val}
}
