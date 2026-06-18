package trees

import "testing"

func TestInsertAndSearch(t *testing.T) {
	tree := &BPtree{}

	pairs := [][2]int{{1, 10}, {2, 20}, {3, 30}, {4, 40}, {5, 50}}
	for _, kv := range pairs {
		tree.Insert(kv[0], kv[1])
	}

	for _, kv := range pairs {
		v, ok := tree.Search(kv[0])
		if !ok {
			t.Errorf("key %d not found", kv[0])
		} else if v != kv[1] {
			t.Errorf("key %d: got %d, want %d", kv[0], v, kv[1])
		}
	}
}

func TestSearchMissing(t *testing.T) {
	tree := &BPtree{}
	tree.Insert(1, 100)
	tree.Insert(2, 200)

	if _, ok := tree.Search(99); ok {
		t.Error("expected missing key 99 to return false")
	}
}

func TestInsertOrder(t *testing.T) {
	// insert in reverse order — tree should still find all keys
	tree := &BPtree{}
	for i := 10; i >= 1; i-- {
		tree.Insert(i, i*10)
	}
	for i := 1; i <= 10; i++ {
		v, ok := tree.Search(i)
		if !ok {
			t.Errorf("key %d not found", i)
		} else if v != i*10 {
			t.Errorf("key %d: got %d, want %d", i, v, i*10)
		}
	}
}

func TestEmptyTree(t *testing.T) {
	tree := &BPtree{}
	if _, ok := tree.Search(1); ok {
		t.Error("expected empty tree search to return false")
	}
}
