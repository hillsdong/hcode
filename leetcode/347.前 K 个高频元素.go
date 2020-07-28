package leetcode

import (
	"container/heap"
	"sort"
)

// 大顶堆 全部入堆
func topKFrequent(nums []int, k int) []int {
	nHash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nHash[nums[i]]++
	}

	nvh := &nvheap{}
	for v, n := range nHash {
		*nvh = append(*nvh, nv{v: v, n: n})
	}

	heap.Init(nvh)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(nvh).(nv).v
	}

	return res
}

type nv struct {
	v int
	n int
}

type nvheap []nv

func (v nvheap) Len() int {
	return len(v)
}

func (v nvheap) Less(i, j int) bool {
	return v[i].n > v[j].n
}
func (v nvheap) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
	return
}
func (v *nvheap) Pop() interface{} {
	ret := (*v)[(*v).Len()-1]
	*v = (*v)[:(*v).Len()-1]
	return ret
}
func (v *nvheap) Push(x interface{}) {
	*v = append(*v, x.(nv))
	return
}

// 小顶堆 k个元素
func topKFrequent5(nums []int, k int) []int {
	nHash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nHash[nums[i]]++
	}

	k1 := k
	nfs := make(nfqs, k1)
	for v, n := range nHash {
		if k1 > 0 {
			nfs[k1-1] = nf{v: v, n: n}
			k1--
			continue
		}

		if k1 == 0 {
			heap.Init(nfs)
			k1--
		}

		if nfs[0].n < n {
			nfs[0] = nf{v: v, n: n}
			heap.Fix(nfs, 0)
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = nfs[i].v
	}

	return res
}

// 桶排
func topKFrequent4(nums []int, k int) []int {
	res := []int{}

	// 数量Hash
	nHash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nHash[nums[i]]++
	}

	// 装桶
	buckets := make([][]int, len(nums)+1)
	for v, n := range nHash {
		if len(buckets[n]) != 0 {
			buckets[n] = append(buckets[n], v)
		} else {
			buckets[n] = []int{v}
		}
	}

	// 出桶
	for i := len(buckets) - 1; i >= 0 && len(res) < k; i-- {
		if len(buckets[i]) == 0 {
			continue
		}
		res = append(res, buckets[i]...)
	}
	return res[0:k]
}

// 快选
func topKFrequent3(nums []int, k int) []int {
	nHash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nHash[nums[i]]++
	}

	nfs := nfqs{}
	for v, n := range nHash {
		nfs = append(nfs, nf{v: v, n: n})
	}

	quickSelect(nfs, k)

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = nfs[len(nfs)-1-i].v
	}

	return res
}

type nf struct {
	v int
	n int
}

type nfqs []nf

func (v nfqs) Len() int {
	return len(v)
}

func (v nfqs) Less(i, j int) bool {
	return v[i].n < v[j].n
}
func (v nfqs) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
	return
}
func (v nfqs) Pop() interface{} {
	return nil
}
func (v nfqs) Push(x interface{}) {
	return
}

// QuickSelectInterface 快选接口
type QuickSelectInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func quickSelect(items QuickSelectInterface, topK int) {
	_quickSelect(items, 0, items.Len()-1, topK)
	return
}

func _quickSelect(items QuickSelectInterface, s int, e int, k int) {
	l := e - s + 1
	if l <= 1 {
		return
	}

	m := s

	for i := s; i < e; i++ {
		if items.Less(i, e) {
			items.Swap(i, m)
			m++
		}
	}
	items.Swap(m, e)

	if items.Len()-k == m {
		return
	} else if items.Len()-k > m {
		_quickSelect(items, m+1, e, k)
	} else {
		_quickSelect(items, s, m-1, k)
	}
}

// 排序
func topKFrequent2(nums []int, k int) []int {
	nHash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		nHash[nums[i]]++
	}

	type nf struct {
		v int
		n int
	}

	nfs := []nf{}
	for v, n := range nHash {
		nfs = append(nfs, nf{v: v, n: n})
	}
	sort.Slice(nfs, func(i int, j int) bool {
		return nfs[i].n < nfs[j].n
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = nfs[len(nfs)-1-i].v
	}

	return res
}
