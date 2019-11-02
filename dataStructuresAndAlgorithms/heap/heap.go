package heap

import (
	"bytes"
	"fmt"
)

func New(n int, isTopMin bool) *Heap {
	data := make([]int, n+1)
	return &Heap{data, 0, n, isTopMin}
}

func NewWithData(data []int, isTopMin bool) *Heap {
	n := len(data)
	h := &Heap{data, n - 1, n - 1, isTopMin}
	for i := n / 2; i > 0; i-- {
		h.heapify(i)
	}
	return h
}

type Heap struct {
	data     []int
	len      int
	cap      int
	isTopMin bool
}

func (h *Heap) isTopper(i, j int) bool {
	if h.isTopMin {
		return i < j
	}
	return i > j
}

func (h *Heap) Insert(d int) bool {
	if h.len == h.cap {
		return false
	}
	h.len++
	h.data[h.len] = d

	i := h.len
	pi := i / 2
	for pi > 0 && h.isTopper(h.data[i], h.data[pi]) {
		h.data[i], h.data[pi] = h.data[pi], h.data[i]
		i = pi
		pi = i / 2
	}
	return true
}

func (h *Heap) heapify(i int) {
	sli := 2 * i
	sri := 2*i + 1
	for {
		topperi := i
		if sli <= h.len && h.isTopper(h.data[sli], h.data[topperi]) {
			topperi = sli
		}
		if sri <= h.len && h.isTopper(h.data[sri], h.data[topperi]) {
			topperi = sri
		}
		if topperi == i {
			break
		}

		h.data[i], h.data[topperi] = h.data[topperi], h.data[i]
		i = topperi
		sli = 2 * i
		sri = 2*i + 1
	}
	return
}

func (h *Heap) RemoveTop() int {
	if h.len == 0 {
		return 0
	}
	top := h.data[1]
	h.data[1] = h.data[h.len]
	h.len--
	h.heapify(1)

	return top
}

func (h *Heap) ReplaceTop(d int) int {
	if h.len == 0 {
		return 0
	}
	top := h.data[1]
	h.data[1] = d
	h.heapify(1)

	return top
}

func (h *Heap) Top() int {
	return h.data[1]
}

func (h *Heap) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Heap: Lenght = %d, ", h.len))

	buffer.WriteString("Data = [")
	for i := 1; i <= h.len; i++ {
		buffer.WriteString(fmt.Sprintf("%d", h.data[i]))
		if i < h.len {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
