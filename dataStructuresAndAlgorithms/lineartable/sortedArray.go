package lineartable

import (
	"bytes"
	"fmt"
)

// SortedArrayInterface 有序数组 接口
type SortedArrayInterface interface {
	Add(int64)
	Remove(int64)
	Find(int64) int
}

// NewSortedArray 创建新的有序数组
func NewSortedArray(cap int) *SortedArray {
	return &SortedArray{
		data: make([]int64, cap, cap),
		cap:  cap,
		len:  0,
	}
}

// SortedArray 有序数组
type SortedArray struct {
	data []int64
	cap  int
	len  int
}

// Add 增加元素
func (a *SortedArray) Add(v int64) {
	if a.len == a.cap {
		panic("array is full!")
	}
	insertI := a.len
	for i := a.len; i > 0; i-- {
		if a.data[i-1] <= v {
			break
		}
		a.data[i] = a.data[i-1]
		insertI = i - 1
	}
	a.data[insertI] = v
	a.len++
}

// Find 查找元素
func (a *SortedArray) Find(v int64) int {
	for i := 0; i < a.len; i++ {
		if v == a.data[i] {
			return i
		}
	}
	return -1
}

// Remove 删除元素
func (a *SortedArray) Remove(v int64) {
	removeI := a.Find(v)
	if removeI != -1 {
		for i := removeI; i < a.len-1; i++ {
			a.data[i] = a.data[i+1]
		}
		a.len--
	}
}

func (a *SortedArray) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("SortedArray: length = %d, capacity = %d, ", a.len, a.cap))
	buffer.WriteString("data = [")
	for i := 0; i < a.len; i++ {
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i < a.len-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
