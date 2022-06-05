package main

import "container/list"

// 使用内置双向链表实现
type entry struct {
	key, value int
}

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	lst   *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*list.Element{}, list.New()}
}

func (c *LRUCache) Get(key int) int {
	e := c.cache[key]
	if e == nil {
		return -1
	}
	c.lst.MoveToFront(e) // 刷新缓存使用时间
	return e.Value.(entry).value
}

func (c *LRUCache) Put(key, value int) {
	if e := c.cache[key]; e != nil {
		e.Value = entry{key, value}
		c.lst.MoveToFront(e) // 刷新缓存使用时间
		return
	}
	c.cache[key] = c.lst.PushFront(entry{key, value})
	if len(c.cache) > c.cap {
		delete(c.cache, c.lst.Remove(c.lst.Back()).(entry).key)
	}
}
