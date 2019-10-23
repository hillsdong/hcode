package lineartable

import (
	"bytes"
	"fmt"
)

// ArrayInterface 动态数组 接口
type ArrayInterface interface {
	Add(int, interface{})
	Remove(int) interface{}
	Find(interface{}) int
	FindAll(interface{}) []int
	Get(int) interface{}
	Set(int, interface{})
}

// NewArray 创建新的动态数组
func NewArray(cap int) *Array {
	return &Array{
		data: make([]interface{}, cap, cap),
		cap:  cap,
		len:  0,
	}
}

// Array 动态数组实现
type Array struct {
	data []interface{}
	cap  int
	len  int
}

func (a *Array) check(i int, s int, e int) {
	if i < s || i > e {
		panic("index over")
	}
}

func (a *Array) resize(cap int) {
	if cap <= a.len {
		panic("resize error")
	}
	a.cap = cap
	data := make([]interface{}, a.cap, a.cap)
	for i := 0; i < a.len; i++ {
		data[i] = a.data[i]
	}
	a.data = data
}

// Add 添加
func (a *Array) Add(i int, v interface{}) {
	a.check(i, 0, a.len)
	if a.len == a.cap {
		a.resize(a.cap * 2)
	}
	for j := a.len - 1; j >= i; j-- {
		a.data[j+1] = a.data[j]
	}
	a.data[i] = v
	a.len++
}

// Remove 删除
func (a *Array) Remove(i int) interface{} {
	a.check(i, 0, a.len-1)
	v := a.data[i]
	for j := i; j < a.len-1; j++ {
		a.data[j] = a.data[j+1]
	}
	a.data[a.len-1] = nil
	a.len--
	return v
}

// Find 查找索引
func (a *Array) Find(v interface{}) int {
	for i := 0; i < a.len; i++ {
		if a.data[i] == v {
			return i
		}
	}
	return -1
}

// FindAll 查找所有
func (a *Array) FindAll(v interface{}) []int {
	is := []int{}
	for i := 0; i < a.len; i++ {
		if a.data[i] == v {
			is = append(is, i)
		}
	}
	return is
}

// Get 查找值
func (a *Array) Get(i int) interface{} {
	a.check(i, 0, a.len-1)
	return a.data[i]
}

// Set 设置值
func (a *Array) Set(i int, v interface{}) {
	a.check(i, 0, a.len-1)
	a.data[i] = v
}

func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: length = %d, capacity = %d, ", a.len, a.cap))
	buffer.WriteString("data = [")
	for i := 0; i < a.len; i++ {
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i < a.len-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(fmt.Sprintf("]"))
	return buffer.String()
}
