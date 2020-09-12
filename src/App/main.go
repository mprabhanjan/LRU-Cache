package main

import (
    lru "LruCache"
    "fmt"
)

func main() {

    lruCache := lru.NewLruCache(5)

    lruCache.Insert("Medha", 64)
    lruCache.Insert("Mahidas", 45)
    lruCache.Insert("Sahana", 39)
    lruCache.Insert("Samhita", 9)
    lruCache.Insert("Samvruta", 4)

    print_lru_cache(lruCache)
    fmt.Println("===========================")
    lruCache.Insert("Ashok", 70)
    print_lru_cache(lruCache)

    _, ok := lruCache.Get("Medha")
    if ok {
        fmt.Printf("Unexpected key in LruCache 'Medha'!!")
    }

    _, _ = lruCache.Get("Mahidas")
    _, _ = lruCache.Get("Sahana")
    _, _ = lruCache.Get("Samhita")
    _, _ = lruCache.Get("Samvruta")

    lruCache.Insert("Medha", 64)

    _, ok = lruCache.Get("Ashok")
    if ok {
        fmt.Printf("Unexpected key in LruCache 'Ashok'!!")
    }

    _, ok = lruCache.Get("Medha")
    if !ok {
        fmt.Printf("Unexpected key 'Medha' not found in LruCache!!")
    }

    lruCache.Insert("Mahidas", 46)
    val, ok := lruCache.Get("Mahidas")
    age, ok := val.(int)
    if ok {
        fmt.Printf("key 'Mahidas' has %d value in LruCache!!\n", age)
    }

    fmt.Println("===========================")
    print_lru_cache(lruCache)
}

func print_lru_cache(cache *lru.LruCache) {
    keys := make([]string, 0)
    vals := make([]interface{}, 0)
    keys, vals = cache.DumpAll(keys, vals)
    for i, key := range keys {
        if val, ok := vals[i].(int); ok {
            fmt.Printf("Key %s, val %d\n", key, val)
        } else {
            fmt.Printf("Type %T\n", vals[i])
        }
    }
}