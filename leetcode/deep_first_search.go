package main

// TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListToBinaryTree converts a slice of integers to a binary tree using level-order insertion.
// The tree is constructed where for index i:
//   - Left child is at index 2*i + 1
//   - Right child is at index 2*i + 2
func ListToBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: nums[0]}

	// Use a queue to build the tree level by level
	queue := []*TreeNode{root}
	i := 1

	for i < len(nums) {
		// Get the current node from the queue
		current := queue[0]
		queue = queue[1:]

		// Assign left child
		if i < len(nums) {
			current.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, current.Left)
			i++
		}

		// Assign right child
		if i < len(nums) {
			current.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, current.Right)
			i++
		}
	}

	return root
}

func inorderTraversal(root *TreeNode) []int {
	nums := []int{}

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		nums = append(nums, root.Val)
		traverse(root.Right)

		return
	}

	traverse(root)

	return nums
}

// func main() {
// 	nums := []int{1, 2, 3}

// 	// Convert list to binary tree
// 	root := ListToBinaryTree(nums)

// 	// Test with inorder traversal
// 	fmt.Println("Input:", nums)

// 	fmt.Println("Inorder traversal:", inorderTraversal(root))
// }
