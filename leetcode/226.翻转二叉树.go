package leetcode

import (
	"bytes"
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	b := bytes.Buffer{}

	// DFS
	//b = t._string(t, b)

	// BFS
	ns := list.New()
	ns.PushBack(t)
	for ns.Len() != 0 {
		n := ns.Front().Value.(*TreeNode)
		b.WriteString(fmt.Sprintf("V:%d ", n.Val))
		if n.Left != nil {
			ns.PushBack(n.Left)
		}
		if n.Right != nil {
			ns.PushBack(n.Right)
		}
		ns.Remove(ns.Front())
	}

	return b.String()
}

func (t *TreeNode) _string(root *TreeNode, b bytes.Buffer) bytes.Buffer {
	b.WriteString(fmt.Sprintf("V:%d ", root.Val))
	if root.Left != nil {
		b = t._string(root.Left, b)
	}
	if root.Right != nil {
		b = t._string(root.Right, b)
	}
	return b
}

func invertTree(root *TreeNode) *TreeNode {
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	root.Left, root.Right = root.Right, root.Left
	return root
}
