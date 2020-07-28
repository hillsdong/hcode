package leetcode

import (
	"bytes"
	"fmt"
)

// TODO 递归

// 思路：不断寻找分组结尾元素，找到后翻转本分组，并与上一分组相连，最后连接剩余元素即可
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 分组辅助变量：上一个分组头，分组头，分组尾
	var kheadpre, khead, ktail *ListNode = nil, head, nil
	// 循环辅助变量，返回变量
	var cur, ret *ListNode = head, head
	// 循环链表
	for i := 1; cur != nil; i++ {
		// 分组结尾元素，要时刻准备着
		ktail, cur = cur, cur.Next

		// 终于到到分组结尾啦
		if i%k == 0 {
			// 赶快翻转组内元素
			_reverseListK(khead, k)

			// 敲黑板！由于分组内元素已经翻转，所以分组头尾已互换
			// 判断是否为首个分组
			if kheadpre != nil {
				// 若不是首个分组，则将上一个分组尾指向本分组头
				kheadpre.Next = ktail
			} else {
				// 第一个分组头作为元素返回
				ret = ktail
			}
			// 重置上一个分组头和分组头
			kheadpre, khead = khead, cur
		}
	}
	// 连接剩余元素
	if kheadpre != nil {
		kheadpre.Next = khead
	}
	// 返回
	return ret
}

// 翻转前K个链表元素
func _reverseListK(head *ListNode, k int) {
	var pre, cur *ListNode = nil, head
	for i := 0; i < k; i++ {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return
}

func (l ListNode) String() string {
	var buffer bytes.Buffer
	ele := &l
	i := 0
	for ele != nil && i < 10 {
		buffer.WriteString(fmt.Sprintf("%v->", ele.Val))
		ele = ele.Next
		i++
	}
	return buffer.String()
}
