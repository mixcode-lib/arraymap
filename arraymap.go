package arraymap

import "sort"

// ArrayMap is a key-value array with map of keys to the array index.
// It works like an ordered map that preserves the inserting order.
type ArrayMap[K comparable, V any] struct {
	Key   []K       // list of keys, in append order.
	Value []V       // list of values, in append order.
	Index map[K]int // index of an entry of given key.
}

// Make a array map.
func NewArrayMap[K comparable, V any]() *ArrayMap[K, V] {
	return &ArrayMap[K, V]{
		Key:   make([]K, 0),
		Value: make([]V, 0),
		Index: make(map[K]int),
	}
}

// The number of entries.
func (m *ArrayMap[K, V]) Len() int {
	return len(m.Key)
}

// Put a key-value pair.
// If the key already exists, then the value is overwritten.
func (m *ArrayMap[K, V]) Put(key K, value V) {
	idx, ok := m.Index[key]
	if ok {
		// key already in the map
		m.Value[idx] = value
	} else {
		// new key value pair
		m.Key = append(m.Key, key)
		m.Value = append(m.Value, value)
		m.Index[key] = len(m.Key) - 1
	}
}

// Put multiple key-value pairs.
func (m *ArrayMap[K, V]) PutValues(key []K, value []V) {
	l, lv := len(key), len(value)
	if l > lv {
		l = lv
	}
	for i := 0; i < l; i++ {
		m.Put(key[i], value[i])
	}
}

// Append a ArrayMap to another ArrayMap.
func (m *ArrayMap[K, V]) Append(m2 *ArrayMap[K, V]) {
	m.PutValues(m2.Key, m2.Value)
}

// Fetch a value with the key.
// The zero T is returned when the key does not exists in the ArrayMap.
func (m *ArrayMap[K, V]) Fetch(key K) (value V) {
	value, _ = m.Get(key)
	return
}

// Get a value with the key, with validity check.
// If the key does not exists then ok will be false.
func (m *ArrayMap[K, V]) Get(key K) (value V, ok bool) {
	idx, ok := m.Index[key]
	if ok {
		value = m.Value[idx]
	}
	return
}

// Get a key-value pair by the index.
func (m *ArrayMap[K, V]) GetAt(index int) (K, V) {
	return m.Key[index], m.Value[index]
}

// Test if the key is in the ArrayMap.
func (m *ArrayMap[K, V]) HasKey(key K) (ok bool) {
	_, ok = m.Index[key]
	return
}

// Delete entries with keys.
// The indexs of other entries may be changed after deletion.
// Also note that this operation is a bit costly because arrays are copied internally.
func (m *ArrayMap[K, V]) Delete(key ...K) {
	// make a list of index to be deleted
	idx := make([]int, 0, len(key)+1)
	for i := range key {
		n, ok := m.Index[key[i]]
		if ok {
			idx = append(idx, n)
		}
	}
	m.DeleteAt(idx...)
}

// Delete entries at indexes.
// The indexs of other entries may be changed after the deletion.
// Note that this operation is a bit costly because arrays are copied internally.
func (m *ArrayMap[K, V]) DeleteAt(idx ...int) {
	if len(idx) == 0 {
		return
	}
	count := len(idx)
	idx = append(idx, len(m.Key))
	sort.Ints(idx)

	w := idx[0] // array write position
	for i := 0; i < count; i++ {
		delete(m.Index, m.Key[idx[i]]) // delete the key
		// pack the array
		l := idx[i+1] - (idx[i] + 1) // length of non-deleting block between delete candidates
		if l == 0 {
			continue
		}
		copy(m.Key[w:], m.Key[idx[i]+1:idx[i+1]])
		copy(m.Value[w:], m.Value[idx[i]+1:idx[i+1]])
		w += l
	}
	m.Key = m.Key[:w]
	m.Value = m.Value[:w]
	// rebuild index of the remaining keys
	for i := range m.Key {
		m.Index[m.Key[i]] = i
	}
}
