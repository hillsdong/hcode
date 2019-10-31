package btree

import (
	"bytes"
	"fmt"
)

func New() *Tree {
	return &Tree{}
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
	len  int
}

func (t *Tree) Insert(v int) {
	if t.root == nil {
		t.root = &Node{data: v}
		t.len++
		return
	}
	n := t.root
	t.insert(n, v)

}

func (t *Tree) insert(n *Node, v int) {
	if v < n.data {
		if n.left == nil {
			n.left = &Node{data: v}
			t.len++
			return
		}
		t.insert(n.left, v)
	} else {
		if n.right == nil {
			n.right = &Node{data: v}
			t.len++
			return
		}
		t.insert(n.right, v)
	}
}

func (t *Tree) Delete(v int) bool {
	pn, n := t.findWithParent(nil, t.root, v)
	var ns *Node
	if n == nil {
		return false
	} else if n.left == nil && n.right == nil {
		ns = nil
	} else if n.left == nil {
		ns = n.right
	} else if n.right == nil {
		ns = n.left
	}

	// 插入子节点
	if pn == nil {
		t.root = ns
	} else {
		if n.data < pn.data {
			pn.left = ns
		} else {
			pn.right = ns
		}
	}
	t.len--
	return true
}

func (t *Tree) findWithParent(pn *Node, n *Node, v int) (pno *Node, cno *Node) {
	if n == nil {
		return pn, nil
	} else if v == n.data {
		return pn, n
	} else if v < n.data {
		return t.findWithParent(n, n.left, v)
	} else {
		return t.findWithParent(n, n.right, v)
	}
	return nil, nil
}

func (t *Tree) Find(v int) *Node {
	return t.find(t.root, v)
}

func (t *Tree) find(n *Node, v int) *Node {
	if n == nil {
		return nil
	} else if v == n.data {
		return n
	} else if v < n.data {
		return t.find(n.left, v)
	} else {
		return t.find(n.right, v)
	}
	return nil
}

func (t *Tree) inOrder(n *Node, do func(n *Node)) {
	if n == nil {
		return
	}
	t.inOrder(n.left, do)
	do(n)
	t.inOrder(n.right, do)
	return
}

func (t *Tree) preOrder(n *Node, do func(n *Node)) {
	if n == nil {
		return
	}
	do(n)
	t.preOrder(n.left, do)
	t.preOrder(n.right, do)
	return
}

func (t *Tree) postOrder(n *Node, do func(n *Node)) {
	if n == nil {
		return
	}
	t.postOrder(n.left, do)
	t.postOrder(n.right, do)
	do(n)
	return
}

func (t *Tree) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("BTree: Lenght = %d, ", t.len))

	bufferWriteString := func(n *Node) {
		buffer.WriteString(fmt.Sprintf("%d, ", n.data))
	}

	buffer.WriteString("PreOrder = [")
	t.preOrder(t.root, bufferWriteString)
	buffer.WriteString("]")

	buffer.WriteString("InOrder = [")
	t.inOrder(t.root, bufferWriteString)
	buffer.WriteString("]")

	buffer.WriteString("PostOrder = [")
	t.postOrder(t.root, bufferWriteString)
	buffer.WriteString("]")

	return buffer.String()
}
