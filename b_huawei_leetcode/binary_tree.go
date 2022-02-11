package b_huawei_leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTreeStack struct {
	array []*TreeNode
}

func (s *BinaryTreeStack) Push(node *TreeNode) {
	s.array = append(s.array, node)
}

func (s *BinaryTreeStack) Pop() *TreeNode {
	if s.Empty() {
		return nil
	}
	node := s.array[s.Len()-1]
	s.array = s.array[:s.Len()-1]
	return node
}

func (s *BinaryTreeStack) Empty() bool {
	return len(s.array) == 0
}

func (s *BinaryTreeStack) Len() int {
	return len(s.array)
}

type BinaryTreeQueue struct {
	array []*TreeNode
}

func (q *BinaryTreeQueue) Push(node *TreeNode) {
	q.array = append(q.array, node)
}

func (q *BinaryTreeQueue) Len() int {
	return len(q.array)
}

func (q *BinaryTreeQueue) Empty() bool {
	return q.Len() == 0
}

func (q *BinaryTreeQueue) Pull() *TreeNode {
	if q.Empty() {
		return nil
	}

	node := q.array[0]
	q.array = q.array[1:]
	return node
}
