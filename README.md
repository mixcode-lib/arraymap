# arraymap

A Go utility package works like an ordered map, i.e. a map that keeps the inserting order.

## Examples
```

func someFunc() {
    amap := arraymap.NewArrayMap[string, int]()

    // add some key and values
    amap.Put("Alice", 101)
    amap.Put("Bob", 102)
    amap.Put("Carol", 103)
    amap.Put("Dave", 104)
    amap.Put("Mallory", 113)
    amap.Put("Sybil", 119)

    // Keys are stored in amap.Key[], and values are stored in amap.Value[].
    // Key and value at the same index matches each other.
    for i, k := range amap.Key {
        v, ok := amap.Get(k) // values can be inferenced with Get()
        if !ok || v != amap.Value[i] {
            panic("invalid key-value")
        }
        fmt.Printf("%s %d\n", k, v)
    }

    // Keys could be deleted.
    amap.Delete("Sybil") // note that indexes of other keys and values are changed

    // Index of keys are stored in amap.Index[] .
    idx := amap.Index["Mallory"]
    amap.DeleteAt(idx)

    // Fetch() is a Get() without validity check
    fmt.Printf("Mallory %d\n", amap.Fetch("Mallory"))
}

```
