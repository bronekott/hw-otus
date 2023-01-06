package hw04lrucache

import (
	"log"
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.Lock()
	defer cache.Unlock()
	if item, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, ok
	}
	return nil, false
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.Lock()
	defer cache.Unlock()

	checkedItem, itemOk := cache.items[key]
	if itemOk {
		checkedItem.Value = cacheItem{key: key, value: value}
		cache.queue.MoveToFront(checkedItem)
		return true
	}
	item := cache.queue.PushFront(cacheItem{key: key, value: value})
	cache.items[key] = item
	if cache.queue.Len() > cache.capacity {
		itemToRemove := cache.queue.Back()
		removedCacheItem, ok := itemToRemove.Value.(cacheItem)
		if !ok {
			log.Println("Failed type assertion")
			return false
		}
		cache.queue.Remove(itemToRemove)
		delete(cache.items, removedCacheItem.key)
	}
	return false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
