package leetcode

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(5)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	obj.Put(5, 5)
	fmt.Println(obj)

	obj.Put(1, 51)
	fmt.Println(obj)

	obj.Put(6, 6)
	fmt.Println(obj)

	fmt.Println(obj.Get(4), obj)
}

func TestLRUCache2(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	fmt.Println(obj)
	obj.Put(2, 2)
	fmt.Println(obj)
	fmt.Println(obj.Get(1))
	fmt.Println(obj)
	obj.Put(3, 3)
	fmt.Println(obj)
	fmt.Println(obj.Get(2))
	fmt.Println(obj)

}
