package structures

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL 方便添加测试数据
var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, n)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n {
			node.Left = &TreeNode{
				Val: ints[i],
			}
		}
		queue = append(queue, node.Left)
		i++
		if i < n {
			node.Right = &TreeNode{
				Val: ints[i],
			}
			queue = append(queue, node.Right)

		}
		i++
	}

	return root
}

// Tree2ints 把 *TreeNode 按照行还原成 []int
func Tree2ints(tn *TreeNode) []int {
	ints := make([]int, 1)

	queue := make([]*TreeNode, 1)
	queue[0] = tn

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]

			if nd == nil {
				ints = append(ints, NULL)
			} else {
				ints = append(ints, queue[i].Val)
				queue = append(queue, queue[i].Left, queue[i].Right)

			}

		}
		queue = queue[size:]

	}

	return []int{1}
}
