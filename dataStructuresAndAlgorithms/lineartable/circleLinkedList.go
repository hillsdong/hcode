package lineartable

import (
	"bytes"
	"errors"
	"fmt"
)

// NewCircleLinkedList 返回链表
func NewCircleLinkedList() *CircleLinkedList {
	h := &CircleListNode{}
	h.Next = h
	return &CircleLinkedList{h, 0}
}

// CircleListNode 节点
type CircleListNode struct {
	Next *CircleListNode
	Data interface{}
}

// CircleLinkedList 链表
type CircleLinkedList struct {
	Head *CircleListNode
	Len  int
}

func (l *CircleLinkedList) findNode(index int) (*CircleListNode, error) {
	if index < 0 || index > l.Len {
		return nil, errors.New("index error")
	}
	p := l.Head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	return p, nil
}

// Insert 插入
func (l *CircleLinkedList) Insert(index int, v interface{}) error {
	if index < 1 || index > l.Len+1 {
		return errors.New("index error")
	}
	p, _ := l.findNode(index - 1)
	//inert
	node := &CircleListNode{p.Next, v}
	p.Next = node
	l.Len++
	return nil
}

// Remove 删除
func (l *CircleLinkedList) Remove(index int) error {
	if index < 1 || index > l.Len {
		return errors.New("index error")
	}
	p, _ := l.findNode(index - 1)
	//remove
	p.Next = p.Next.Next
	l.Len--
	return nil
}

// Get 获取
func (l *CircleLinkedList) Get(index int) (interface{}, error) {
	if index < 1 || index > l.Len {
		return nil, errors.New("index error")
	}
	p, _ := l.findNode(index)
	//get
	return p.Data, nil
}

// Set 修改
func (l *CircleLinkedList) Set(index int, v interface{}) error {
	if index < 1 || index > l.Len {
		return errors.New("index error")
	}
	p, _ := l.findNode(index)
	//set
	p.Data = v
	return nil
}

// Find 查找
func (l *CircleLinkedList) Find(v interface{}) int {
	p := l.Head.Next
	i := 1
	for p.Next != l.Head {
		if p.Data == v {
			return i
		}
		p = p.Next
		i++
	}
	return -1
}

// FindAll 查找所有
func (l *CircleLinkedList) FindAll(v interface{}) (indexs []int) {
	p := l.Head.Next
	i := 1
	for p.Next != l.Head {
		if p.Data == v {
			indexs = append(indexs, i)
		}
		p = p.Next
		i++
	}
	return indexs
}

func (l *CircleLinkedList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("LikedList: Length = %d, ", l.Len))
	buffer.WriteString(fmt.Sprintf("header = [%p,%d,%p->], ", l.Head, l.Head.Data, l.Head.Next))

	node := l.Head.Next
	buffer.WriteString("data = [")
	for node.Next != l.Head {
		buffer.WriteString(fmt.Sprintf("%p,%d,%p->, ", node, node.Data, node.Next))
		node = node.Next
	}
	buffer.WriteString(fmt.Sprintf("%p,%d,%p]", node, node.Data, node.Next))
	return buffer.String()
}
