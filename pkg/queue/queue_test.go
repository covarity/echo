package queue

import "testing"
import "container/list"
import "math/rand"

func ensureEmpty(t *testing.T, q *Queue) {
	if l := q.Len(); l != 0 {
		t.Errorf("q.Len() = %d, want %d", l, 0)
	}
	if e := q.Front(); e.Value != nil {
		t.Errorf("q.Front() = %v, want %v", e, nil)
	}
	if e := q.Back(); e.Value != nil {
		t.Errorf("q.Back() = %v, want %v", e, nil)
	}
	if e := q.PopFront(); e.Value != nil {
		t.Errorf("q.PopFront() = %v, want %v", e, nil)
	}
	if e := q.PopBack(); e.Value != nil {
		t.Errorf("q.PopBack() = %v, want %v", e, nil)
	}
}

func TestNew(t *testing.T) {
	q := New()
	ensureEmpty(t, q)
}

func ensureSingleton(t *testing.T, q *Queue) {
	if l := q.Len(); l != 1 {
		t.Errorf("q.Len() = %d, want %d", l, 1)
	}
	if e := q.Front(); e.Value != 42 {
		t.Errorf("q.Front() = %v, want %v", e, 42)
	}
	if e := q.Back(); e.Value != 42 {
		t.Errorf("q.Back() = %v, want %v", e, 42)
	}
}

func TestSingleton(t *testing.T) {
	q := New()
	ensureEmpty(t, q)
	q.PushFront(Item{Value: 42, Priority: 0})
	ensureSingleton(t, q)
	q.PopFront()
	ensureEmpty(t, q)
	q.PushBack(Item{Value: 42, Priority: 0})
	ensureSingleton(t, q)
	q.PopBack()
	ensureEmpty(t, q)
	q.PushFront(Item{Value: 42, Priority: 0})
	ensureSingleton(t, q)
	q.PopBack()
	ensureEmpty(t, q)
	q.PushBack(Item{Value: 42, Priority: 0})
	ensureSingleton(t, q)
	q.PopFront()
	ensureEmpty(t, q)
}

func TestDuos(t *testing.T) {
	q := New()
	ensureEmpty(t, q)
	q.PushFront(Item{Value: 42, Priority: 0})
	ensureSingleton(t, q)
	q.PushBack(Item{Value: 43, Priority: 0})
	if l := q.Len(); l != 2 {
		t.Errorf("q.Len() = %d, want %d", l, 2)
	}
	if e := q.Front(); e.Value != 42 {
		t.Errorf("q.Front() = %v, want %v", e, 42)
	}
	if e := q.Back(); e.Value != 43 {
		t.Errorf("q.Back() = %v, want %v", e, 43)
	}
}

func ensureLength(t *testing.T, q *Queue, len int) {
	if l := q.Len(); l != len {
		t.Errorf("q.Len() = %d, want %d", l, len)
	}
}

func TestZeroValue(t *testing.T) {
	var q Queue
	q.PushFront(Item{Value: 1, Priority: 0})
	ensureLength(t, &q, 1)
	q.PushFront(Item{Value: 2, Priority: 0})
	ensureLength(t, &q, 2)
	q.PushFront(Item{Value: 3, Priority: 0})
	ensureLength(t, &q, 3)
	q.PushFront(Item{Value: 4, Priority: 0})
	ensureLength(t, &q, 4)
	q.PushFront(Item{Value: 5, Priority: 0})
	ensureLength(t, &q, 5)
	q.PushBack(Item{Value: 6, Priority: 0})
	ensureLength(t, &q, 6)
	q.PushBack(Item{Value: 7, Priority: 0})
	ensureLength(t, &q, 7)
	q.PushBack(Item{Value: 8, Priority: 0})
	ensureLength(t, &q, 8)
	q.PushBack(Item{Value: 9, Priority: 0})
	ensureLength(t, &q, 9)
	const want = "[{5 0} {4 0} {3 0} {2 0} {1 0} {6 0} {7 0} {8 0} {9 0}]"
	if s := q.String(); s != want {
		t.Errorf("q.String() = %s, want %s", s, want)
	}
}

func TestGrowShrink1(t *testing.T) {
	var q Queue
	for i := 0; i < size; i++ {
		q.PushBack(Item{Value: i, Priority: 0})
		ensureLength(t, &q, i+1)
	}
	for i := 0; q.Len() > 0; i++ {
		x := q.PopFront()
		if x.Value != i {
			t.Errorf("q.PopFront() = %d, want %d", x, i)
		}
		ensureLength(t, &q, size-i-1)
	}
}
func TestGrowShrink2(t *testing.T) {
	var q Queue
	for i := 0; i < size; i++ {
		q.PushFront(Item{Value: i, Priority: 0})
		ensureLength(t, &q, i+1)
	}
	for i := 0; q.Len() > 0; i++ {
		x := q.PopBack()
		if x.Value != i {
			t.Errorf("q.PopBack() = %d, want %d", x, i)
		}
		ensureLength(t, &q, size-i-1)
	}
}

const size = 1024

func BenchmarkPushFrontQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue
		for n := 0; n < size; n++ {
			q.PushFront(Item{Value: n, Priority: 0})
		}
	}
}
func BenchmarkPushFrontList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushFront(Item{Value: n, Priority: 0})
		}
	}
}

func BenchmarkPushBackQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue
		for n := 0; n < size; n++ {
			q.PushBack(Item{Value: n, Priority: 0})
		}
	}
}
func BenchmarkPushBackList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushBack(Item{Value: n, Priority: 0})
		}
	}
}
func BenchmarkPushBackChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := make(chan Item, size)
		for n := 0; n < size; n++ {
			q <- Item{Value: n, Priority: 0}
		}
		close(q)
	}
}

var rands []float32

func makeRands() {
	if rands != nil {
		return
	}
	rand.Seed(64738)
	for i := 0; i < 4*size; i++ {
		rands = append(rands, rand.Float32())
	}
}
func BenchmarkRandomQueue(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var q Queue
		for n := 0; n < 4*size; n += 4 {
			if rands[n] < 0.8 {
				q.PushBack(Item{Value: n, Priority: 0})
			}
			if rands[n+1] < 0.8 {
				q.PushFront(Item{Value: n, Priority: 0})
			}
			if rands[n+2] < 0.5 {
				q.PopFront()
			}
			if rands[n+3] < 0.5 {
				q.PopBack()
			}
		}
	}
}
func BenchmarkRandomList(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < 4*size; n += 4 {
			if rands[n] < 0.8 {
				q.PushBack(Item{Value: n, Priority: 0})
			}
			if rands[n+1] < 0.8 {
				q.PushFront(Item{Value: n, Priority: 0})
			}
			if rands[n+2] < 0.5 {
				if e := q.Front(); e != nil {
					q.Remove(e)
				}
			}
			if rands[n+3] < 0.5 {
				if e := q.Back(); e != nil {
					q.Remove(e)
				}
			}
		}
	}
}

func BenchmarkGrowShrinkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue
		for n := 0; n < size; n++ {
			q.PushBack(Item{Value: i, Priority: 0})
		}
		for n := 0; n < size; n++ {
			q.PopFront()
		}
	}
}
func BenchmarkGrowShrinkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushBack(Item{Value: i, Priority: 0})
		}
		for n := 0; n < size; n++ {
			if e := q.Front(); e != nil {
				q.Remove(e)
			}
		}
	}
}
