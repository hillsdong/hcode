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
	// 寻找待删除节点
	pn, n := t.findWithParent(nil, t.root, v)
	if n == nil {
		return false
	}

	// 删除节点拥有左右子节点
	if n.left != nil && n.right != nil {
		minPn, minN := t.findMinWithParent(n, n.right)
		n.data = minN.data
		pn = minPn
		n = minN
	}

	// 寻找待连接节点
	var ns *Node
	if n.left == nil && n.right == nil {
		ns = nil
	} else if n.left == nil {
		ns = n.right
	} else if n.right == nil {
		ns = n.left
	}

	// 连接节点
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
	}

	if v == n.data {
		return pn, n
	}

	pn = n
	if v < n.data {
		n = n.left
	} else {
		n = n.right
	}

	return t.findWithParent(pn, n, v)
}

func (t *Tree) findMinWithParent(pn *Node, n *Node) (pno *Node, cno *Node) {
	if n.left == nil {
		return pn, n
	}

	return t.findMinWithParent(n, n.left)
}

func (t *Tree) Find(v int) *Node {
	return t.find(t.root, v)
}

func (t *Tree) find(n *Node, v int) *Node {
	if n == nil {
		return nil
	}
	if v == n.data {
		return n
	}

	if v < n.data {
		n = n.left
	} else {
		n = n.right
	}

	return t.find(n, v)
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

func (t *Tree) floorOrder(n *Node, do func(n *Node)) {
	ns := []*Node{n}
	fLen, nfLen, fNum := len(ns), 0, 0

	for i := 0; i < len(ns); i++ {
		if ns[i] == nil {
			break
		}
		if nfLen == 0 {
			fmt.Printf("\n n%d: ", fNum+1)
		}
		do(ns[i])
		fmt.Printf("%d ", ns[i].data)
		fLen--
		if ns[i].left != nil {
			ns = append(ns, ns[i].left)
			nfLen++
		}
		if ns[i].right != nil {
			ns = append(ns, ns[i].right)
			nfLen++
		}
		if fLen == 0 {
			fNum++
			fLen, nfLen = nfLen, 0
		}
	}
	fmt.Println("")
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

	buffer.WriteString("FloorOrder = [")
	t.floorOrder(t.root, bufferWriteString)
	buffer.WriteString("]")

	return buffer.String()
}
