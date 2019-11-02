package btree2

import "testing"

// 1
//   3
// 2   4
//       5

func TestInsert(t *testing.T) {
	b := New()
	t.Log(b.Delete(5))
	t.Log(b)
	b.Insert(1)
	t.Log(b.Delete(11))
	t.Log(b)
	t.Log(b.Delete(1))
	t.Log(b)
	b.Insert(3)
	t.Log(b)
	b.Insert(4)
	b.Insert(2)
	b.Insert(5)
	t.Log(b)

	t.Log(b.Find(3))
	t.Log(b.Find(30))

	t.Log(b.Delete(4))
	t.Log(b)

	t.Log(b.Delete(5))
	t.Log(b)

	b.Insert(1)
	b.Insert(5)
	b.Insert(4)
	t.Log(b)

	//    3
	//  2   5
	//1   4

	t.Log(b.Delete(3))
	t.Log(b)

	//    4
	//  2   5
	//1

}
