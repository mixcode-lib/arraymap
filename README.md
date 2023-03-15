# golib-arraymap

A Go utility structure works like an ordered map, i.e. a map that keeps the inserting order.

## Examples
```
    amap := arraymap.NewArrayMap[string, int]()

    // add some key and values
    amap.Set("Alice", 101)
    amap.Set("Bob", 102)
    amap.Set("Carol", 103)
    amap.Set("Dave", 104)
    amap.Set("Mallory", 113)
    amap.Set("Sybil", 119)

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
```
