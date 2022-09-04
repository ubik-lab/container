package list

import (
	"container/list"
	"testing"
)

func work_stdList_any() {
	l := list.New()

	for i := 0; i < 10_000; i++ {
		l.PushBack(i)
		l.PushFront([]int{i})
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(int); ok {
			e.Value = v + 1
		} else {
			l.Remove(e)
		}
	}

	for i := 0; i < 10_000; i++ {
		l.PushBack([]int{i})
		l.PushFront(i)
	}

	// clear
	l.Init()
}

func work_genericList_any() {
	l := New[any]()

	for i := 0; i < 10_000; i++ {
		l.PushBack(i)
		l.PushFront([]int{i})
	}

	for e := l.Front(); e != nil; e = e.Next() {
		if v, ok := e.Value.(int); ok {
			e.Value = v + 1
		} else {
			l.Remove(e)
		}
	}

	for i := 0; i < 10_000; i++ {
		l.PushBack([]int{i})
		l.PushFront(i)
	}

	// clear
	l.Init()
}

func BenchmarkStdList_Any(b *testing.B) {
	for i := 0; i < b.N; i++ {
		work_stdList_any()
	}
}

func BenchmarkGenericList_Any(b *testing.B) {
	for i := 0; i < b.N; i++ {
		work_genericList_any()
	}
}

func work_stdList_int() {
	l := list.New()

	for i := 0; i < 10_000; i++ {
		l.PushBack(i)
		l.PushFront(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		v := e.Value.(int)
		e.Value = v + 1
	}

	for i := 0; i < 10_000; i++ {
		l.PushBack(i)
		l.PushFront(i)
	}

	// clear
	l.Init()
}

func work_genericList_int() {
	l := New[int]()

	for i := 0; i < 10_000; i++ {
		l.PushBack(i)
		l.PushFront(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		e.Value += 1
	}

	for i := 0; i < 10_000; i++ {
		l.PushBack(i)
		l.PushFront(i)
	}

	// clear
	l.Init()
}

func BenchmarkStdList_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		work_stdList_int()
	}
}

func BenchmarkGenericList_Int(b *testing.B) {
	for i := 0; i < b.N; i++ {
		work_genericList_int()
	}
}
