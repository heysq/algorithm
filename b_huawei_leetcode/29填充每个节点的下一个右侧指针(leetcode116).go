package b_huawei_leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}

	connectNode(root.Left, root.Right)
	return root
}

func connectNode(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	connectNode(node1.Left, node2.Right)
	connectNode(node2.Left, node2.Right)
	connectNode(node1.Right, node2.Left)
}
