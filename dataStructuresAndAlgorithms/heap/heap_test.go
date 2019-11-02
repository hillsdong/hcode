package heap

import "testing"

func TestOperate(t *testing.T) {

	h := New(10, false)
	t.Log(h.Top())
	h.Insert(1)
	t.Log(h.Top())
	h.Insert(3)
	t.Log(h.Top())
	h.Insert(2)
	t.Log(h.Top())
	h.Insert(4)
	t.Log(h.Top())
	h.Insert(7)
	t.Log(h.Top())
	h.RemoveTop()
	t.Log(h.Top())
	h.RemoveTop()
	t.Log(h.Top())

}

func TestTopK(t *testing.T) {
	d := []int{2, 3, 4, 1, 9, 8, 4, 3, 7, 8}
	k := 3
	h := New(k, true)

	for i := 0; i < k; i++ {
		h.Insert(d[i])
	}

	for i := k; i < len(d); i++ {
		if d[i] > h.Top() {
			h.ReplaceTop(d[i])
		}
	}

	t.Log(h)
}

func TestMerge(t *testing.T) {
	// sorce datas
	ds := [][]int{[]int{1, 2, 3, 4, 5, 6}, []int{10, 20, 30, 40, 50, 60}, []int{11, 12, 13, 14, 15, 16}}
	dsLen := len(ds)
	var dLen int
	for i := 0; i < dsLen; i++ {
		dLen = dLen + len(ds[i])
	}

	// heap
	h := New(dsLen, true)
	// result data
	d := make([]int, dLen)

	// compute
	dsIHash := make(map[int]int)
	dsPHash := make([]int, dsLen)
	for i := 0; i < dsLen; i++ {
		dsIHash[ds[i][0]] = i
		h.Insert(ds[i][0])
		dsPHash[i]++
	}

	k := 0
	for ; k < dLen-dsLen; k++ {
		d[k] = h.Top()
		//find next
		dsI := dsIHash[d[k]]
		dsP := dsPHash[dsI]
		if dsP == len(ds[dsI]) {
			for i := 0; i < dsLen; i++ {
				if dsPHash[i] < len(ds[i]) {
					dsI = i
					dsP = dsPHash[i]
					break
				}
			}
		}
		data := ds[dsI][dsP]
		dsIHash[data] = dsI
		h.ReplaceTop(data)
		dsPHash[dsI]++
	}

	for i := 0; i < dsLen; i++ {
		d[k+i] = h.RemoveTop()
	}
	t.Log(d)
}

func TestSort(t *testing.T) {
	d := []int{0, 2, 3, 4, 1, 9, 8, 4, 3, 7, 8}
	h := NewWithData(d, false)
	for i := 1; i < len(d); i++ {
		t.Log(h.RemoveTop())
	}
}
