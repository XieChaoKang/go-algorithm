package three

// 迭代方式实现树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代思路如下：从左子树最深的节点开始往上迭代遍历 自底向上
// 1，遍历左子树 直达最深的那个左子树节点 按序压入栈中
// 2，当前root节点的全部左子树节点已经压入栈 开始弹出节点 当前root节点更换为栈弹出的节点
// 3，因为当前root已经是最深左节点，所以当前root节点已经没有左节点了，值Val压入结果集合中，从栈移除该节点，同时root指向其右子节点 左中右的顺序
// 4，重复以上步骤，直到root节点为空且栈内无节点存在
func inorderTraversal(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		// 遍历左子树 直达最深的那个左子树节点 按序压入栈中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 栈顶开始往外弹 进行遍历
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// root指向其右子树 遍历右边
		root = root.Right
	}
	return res
}
