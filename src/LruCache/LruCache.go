package LruCache

import (
    "dll"
)

type cache_object struct {
    key string
    data interface {}
    link *dll.Dlist_link
}

type LruCache struct {
    cache_map map[string]*cache_object
    dll *dll.Dlist
    max_size int

    num_nodes int
    num_inserts int
    num_removes int
    num_gets int
}

func NewLruCache(max_size int) *LruCache {
    if max_size < 1{
        return nil
    }
    new_cache := &LruCache{
        cache_map: make(map[string]*cache_object),
        dll: dll.NewDlist(),
        max_size: max_size,
    }
    return new_cache
}

func (cache *LruCache) Insert(key string, value interface{}) {

    if obj, ok := cache.cache_map[key]; ok {
        obj.data = value
        cache.dll.ReInsertAtHead(obj.link)
        return
    }

    if len(cache.cache_map) >= cache.max_size {
        elem := cache.dll.PopTail()
        obj, ok := elem.(*cache_object)
        if ok {
            delete(cache.cache_map, obj.key)
        }
    }

    obj := &cache_object{
        key: key,
        data: value,
    }
    obj.link = cache.dll.InsertAtHead(obj)
    cache.cache_map[key] = obj
    return
}

func (cache *LruCache) Get(key string) (interface{}, bool) {

    obj, ok := cache.cache_map[key]
    if ok {
        cache.dll.ReInsertAtHead(obj.link)
        return obj.data, true
    }

    return nil, false
}

func (cache *LruCache) Remove(key string) interface {} {
    obj, ok := cache.cache_map[key]
    if !ok {
        return nil
    }

    cache.dll.RemoveLink(obj.link)
    delete(cache.cache_map, key)
    obj.link = nil
    return obj.data
}

func (cache *LruCache) DumpAll(keys []string, vals[]interface{}) ([]string, []interface{}){

    for k, val := range cache.cache_map {
        keys = append(keys, k)
        vals = append(vals, val.data)
    }
    return keys, vals
}