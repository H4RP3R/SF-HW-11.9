package queue

import (
	"errors"
	"testing"
)

func compareElems[T comparable](q Queue[T], elems []T) bool {
	if q.Size() != len(elems) {
		return false
	}

	idx := 0
	for e := q.storage.Front(); e != nil; e = e.Next() {
		if e.Value != elems[idx] {
			return false
		}
		idx++
	}

	return true
}

func TestQueueAddInt(t *testing.T) {
	type want struct {
		size  int
		elems []int
	}

	var tests = []struct {
		name string
		item int
		want
	}{
		{"Add int 1", 1, want{size: 1, elems: []int{1}}},
		{"Add int 2", 2, want{size: 2, elems: []int{1, 2}}},
		{"Add int 3", 3, want{size: 3, elems: []int{1, 2, 3}}},
	}

	queue := NewQueue[int]()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue.Add(tt.item)
			equal := compareElems(queue, tt.elems)
			if !equal {
				t.Errorf("elements do not match")
			}
			qSize := queue.Size()
			if qSize != tt.want.size {
				t.Errorf("got size: %d, want size %d", qSize, tt.want.size)
			}
		})
	}
}

func TestQueuePollString(t *testing.T) {
	type want struct {
		el   string
		size int
		err  error
	}

	var tests = []struct {
		name string
		want
	}{
		{"Poll 1", want{el: "a", size: 2, err: nil}},
		{"Poll 2", want{el: "bb", size: 1, err: nil}},
		{"Poll 3", want{el: "ccc", size: 0, err: nil}},
		{"Poll 4", want{el: "", size: 0, err: ErrPollFromEmptyQueue}},
	}
	queue := NewQueue[string]()
	queue.Add("a")
	queue.Add("bb")
	queue.Add("ccc")
	for _, tt := range tests {
		el, err := queue.Poll()
		size := queue.Size()
		if el != tt.want.el {
			t.Errorf("got el: %s, want el: %s", el, tt.want.el)
		}
		if !errors.Is(err, tt.want.err) {
			t.Errorf("got err: %v, want err: %v", err, ErrPollFromEmptyQueue)
		}
		if size != tt.want.size {
			t.Errorf("got size: %d, want size: %d", size, tt.want.size)
		}
	}
}

func TestQueueRemoveString(t *testing.T) {
	type want struct {
		elems []string
		size  int
		ok    bool
	}

	var tests = []struct {
		name string
		el   string
		want
	}{
		{"Remove 1", "xxx", want{elems: []string{"a", "bb", "y", "ccc"}, size: 4, ok: true}},
		{"Remove 2", "y", want{elems: []string{"a", "bb", "ccc"}, size: 3, ok: true}},
		{"Remove 3", "a", want{elems: []string{"bb", "ccc"}, size: 2, ok: true}},
		{"Remove 4", "123", want{elems: []string{"bb", "ccc"}, size: 2, ok: false}},
		{"Remove 5", "ccc", want{elems: []string{"bb"}, size: 1, ok: true}},
		{"Remove 6", "bb", want{elems: []string{}, size: 0, ok: true}},
		{"Remove 7", "a", want{elems: []string{}, size: 0, ok: false}},
	}

	queue := NewQueue[string]()
	queue.Add("a")
	queue.Add("xxx")
	queue.Add("bb")
	queue.Add("y")
	queue.Add("ccc")
	for _, tt := range tests {
		ok := queue.Remove(tt.el)
		size := queue.Size()
		equal := compareElems(queue, tt.elems)
		if !equal {
			t.Errorf("elements do not match")
		}
		if ok != tt.want.ok {
			t.Errorf("got ok: %t, want ok: %t", ok, tt.want.ok)
		}
		if size != tt.want.size {
			t.Errorf("got size: %d, want size: %d", size, tt.want.size)
		}
	}
}
