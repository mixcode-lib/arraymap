package arraymap

import (
	"testing"
)

func TestArrayMap(t *testing.T) {

	// test values
	type T struct {
		K int
		V string
	}
	t1 := []T{{10, ""}, {11, "a"}, {12, "b"}, {13, "c"}, {14, "d"}, {15, "e"}, {16, "f"}}
	ka, va := make([]int, len(t1)), make([]string, len(t1))
	for i := range t1 {
		ka[i], va[i] = t1[i].K, t1[i].V
	}

	am := NewArrayMap[int, string]()

	for _, e := range t1 {
		am.Set(e.K, e.V)
	}
	if am.Len() != len(t1) {
		t.Fatalf("Len() fail")
	}

	// GetAt()
	for i := 0; i < am.Len(); i++ {
		k, v := am.GetAt(i)
		if k != t1[i].K || v != t1[i].V {
			t.Fatalf("GetAt() fail")
		}
	}
	// index out-of-order; must panic
	mustPanic := func(f func()) (panicked bool) {
		defer func() {
			if r := recover(); r != nil {
				panicked = true
			}
		}()
		f()
		return false
	}
	if !mustPanic(func() { am.GetAt(-1) }) {
		t.Fatalf("function did not fail at invalid index")
	}
	if !mustPanic(func() { am.GetAt(len(t1)) }) {
		t.Fatalf("function did not fail at invalid index")
	}

	// Get()
	for i := 0; i < len(t1); i++ {
		v := am.Fetch(t1[i].K)
		if v != t1[i].V {
			t.Fatalf("Fetch() fail")
		}
		v, ok := am.Get(t1[i].K)
		if v != t1[i].V || !ok {
			t.Fatalf("Get() fail")
		}
		if !am.HasKey(t1[i].K) {
			t.Fatalf("HasKey() fail")
		}
	}
	v, ok := am.Get(-1)
	if v != "" || ok {
		t.Fatalf("Get() fail")
	}
	if am.HasKey(-1) {
		t.Fatalf("HasKey() fail")
	}

	// Delete()
	keyList := make([]int, 0)
	for i := 0; i < len(t1)/2; i += 2 {
		keyList = append(keyList, t1[i].K)
	}
	am.Delete(keyList...)
	for i := 0; i < len(t1)/2; i += 2 {
		_, ok := am.Get(t1[i].K)
		if ok {
			t.Fatalf("Delete() fail")
		}
		_, ok = am.Get(t1[i+1].K)
		if !ok {
			t.Fatalf("Delete() fail")
		}
	}

	// DeleteAt()
	keyList = keyList[:0]
	for i := 1; i < am.Len()-1; i++ {
		keyList = append(keyList, am.Key[i])
	}
	am.DeleteAt(0, am.Len()-1)
	if am.Len() != len(keyList) {
		t.Fatalf("DeleteAt() fail")
	}
	for i := range keyList {
		_, ok := am.Get(keyList[i])
		if !ok {
			t.Fatalf("DeleteAt() fail")
		}
	}

}
