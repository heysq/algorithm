package b_huawei_leetcode

import "math"

// 计算从起点 start 到终点 target 的最近距离
// int BFS(Node start, Node target) {
//     Queue<Node> q; // 核心数据结构
//     Set<Node> visited; // 避免走回头路

//     q.offer(start); // 将起点加入队列
//     visited.add(start);
//     int step = 0; // 记录扩散的步数

//     while (q not empty) {
//         int sz = q.size();
//         /* 将当前队列中的所有节点向四周扩散 */
//         for (int i = 0; i < sz; i++) {
//             Node cur = q.poll();
//             /* 划重点：这里判断是否到达终点 */
//             if (cur is target)
//                 return step;
//             /* 将 cur 的相邻节点加入队列 */
//             for (Node x : cur.adj()) {
//                 if (x not in visited) {
//                     q.offer(x);
//                     visited.add(x);
//                 }
//             }
//         }
//         /* 划重点：更新步数在这里 */
//         step++;
//     }
// }

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归算法
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var md = math.MaxInt
	if root.Left != nil {
		md = min(minDepth(root.Left), md)
	}
	if root.Right != nil {
		md = min(minDepth(root.Right), md)
	}
	return md + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	var depth int = 1
	queue := &BinaryTreeQueue{}
	queue.Push(root)
	for !queue.Empty() {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Pull()
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue.Push(node.Left)
			}
			if node.Right != nil {
				queue.Push(node.Right)
			}
		}
		depth++
	}
	return depth
}
