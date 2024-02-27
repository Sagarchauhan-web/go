package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue length: %v", len(q.items))
		}

		if !q.Add(i) {
			t.Errorf("Append failed: %v", i)
		}
	}

	if q.Add(4) {
		t.Errorf("Append should not succeed:")
	}
}

func TextNext(t *testing.T) {
	q := New(3)
	for i := 0; i< 3; i++ {
		q.Add(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()

		if !ok {
			t.Errorf("should be able to get item from queue")
		}

		if item != i {
			t.Errorf("got item in wrong order")
		}
	}

	item, ok := q.Next()
	if ok {
		t.Errorf("Should not be any more items in queue, got: %v", item)
	}
}