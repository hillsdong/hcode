package lineartable

import (
	"bytes"
	"errors"
	"fmt"
)

// LinkedListInterface 单链表 接口
type LinkedListInterface interface {
	Insert(int, interface{}) error
	Remove(int) error
	Get(int) (interface{}, error)
	Set(int, interface{}) error
	Find(interface{}) int
	FindAll(interface{}) []int
}

// NewLinkedList 返回链表
func NewLinkedList() *LinkedList {
	return &LinkedList{&ListNode{nil, 0}, 0}
}

// ListNode 节点
type ListNode struct {
	Next *ListNode
	Data interface{}
}

// LinkedList 链表
type LinkedList struct{
	Head *ListNode
	Len  int
}

// Reverse reverses the list
func (l *LinkedList) Reverse(){
	if l.Len < 1 {
		return
	}
	node := l.Head
	nodeNext := node.Next
	for nodeNext != nil {
		nodeNextNext := nodeNext.Next
		nodeNext.Next = node

		node = nodeNext
		nodeNext = nodeNextNext
	}
	l.Head.Next.Next = nil
	l.Head.Next = node
	return
}

// FindMidNode return the mid node of the list
func (l *LinkedList) FindMidNode() *ListNode{
	if l.Len == 0 {
		return nil
	}
	node := l.Head
	for i:=0; i<=l.Len/2; i++ {
		node = node.Next
	}
	return node
}

// FindMidNodeWithoutLen return the mid node of the list, but do not use the list's lenght
func (l *LinkedList) FindMidNodeWithoutLen() *ListNode{
	if l.Head.Next == nil {
		return nil
	}
	if l.Head.Next.Next == nil{
		return l.Head.Next
	}
	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func (l *LinkedList) findNode(index int) (*ListNode, error) {
	if index < 0 || index > l.Len {
		return nil, errors.New("index error")
	}
	p := l.Head
	for i:=0; i < index; i++ {
		p = p.Next
	}
	return p, nil
}


// Insert 插入
func (l *LinkedList) Insert(index int, v interface{}) error {
	if index < 1 || index > l.Len + 1 {
		return errors.New("index error")
	}
	p,_ := l.findNode(index-1)
	//inert
	node := &ListNode{p.Next, v}
	p.Next = node
	l.Len++
	return nil
}

// Remove 删除
func (l *LinkedList) Remove(index int) error {
	if index < 1 || index > l.Len {
		return errors.New("index error")
	}
	p,_ := l.findNode(index-1)
	//remove
	p.Next = p.Next.Next
	l.Len--
	return nil
}

// Get 获取
func (l *LinkedList) Get(index int) (interface{}, error) {
	if index < 1 || index > l.Len {
		return nil, errors.New("index error")
	}
	p,_ := l.findNode(index)
	//get
	return p.Data, nil
}

// Set 修改
func (l *LinkedList) Set(index int, v interface{}) error {
	if index < 1 || index > l.Len {
		return errors.New("index error")
	}
	p,_ := l.findNode(index)
	//set
	p.Data = v
	return nil
}

// Find 查找
func (l *LinkedList) Find(v interface{}) int  {
	p := l.Head
	i := 0
	for p.Next != nil {
		p = p.Next
		i++
		if p.Data == v {
			return i
		}
	}
	return -1
}

// FindAll 查找所有
func (l *LinkedList) FindAll(v interface{}) (indexs []int) {
	p := l.Head
	i := 0
	for p.Next != nil {
		p = p.Next
		i++
		if p.Data == v {
			indexs = append(indexs, i)
		}
	}
	return indexs
}

func (l *LinkedList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("LikedList: Length = %d, ", l.Len))
	node := l.Head
	buffer.WriteString("data = [")
	for node.Next != nil {
		buffer.WriteString(fmt.Sprintf("%d->", node.Next.Data))
		node = node.Next
	}
	buffer.WriteString("]")
	return buffer.String()
}
