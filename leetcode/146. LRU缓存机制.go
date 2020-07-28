package leetcode

import (
	"bytes"
	"container/list"
	"fmt"
)

//TODO 自实现双向链表

type kv struct {
	k, v int
}

// LRUCache 最近最少使用缓存
type LRUCache struct {
	hash map[int]*list.Element
	data *list.List
	len  int
	cap  int
}

// Constructor 初始化
func Constructor(capacity int) LRUCache {
	return LRUCache{hash: map[int]*list.Element{}, data: list.New(), cap: capacity}
}

// Get 获取元素
func (l *LRUCache) Get(key int) int {
	if ele, exist := l.hash[key]; exist {
		l.data.MoveToFront(ele)
		l.hash[key] = l.data.Front()
		return ele.Value.(kv).v
	}
	return -1
}

// Put 插入元素
func (l *LRUCache) Put(key int, value int) {
	if ele, exist := l.hash[key]; exist {
		ele.Value = kv{key, value}
		l.data.MoveToFront(ele)
		return
	}

	// 删除最久
	if l.len == l.cap {
		back := l.data.Back()
		delete(l.hash, back.Value.(kv).k)
		l.data.Remove(back)
		l.len--
	}

	// 插入
	l.hash[key] = l.data.PushFront(kv{key, value})
	l.len++
	return
}

func (l LRUCache) String() string {
	var buffer bytes.Buffer
	ele := l.data.Front()
	for ele != nil {
		buffer.WriteString(fmt.Sprintf("%v->", ele.Value))
		ele = ele.Next()
	}
	return buffer.String()
}
