package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type entry struct {
	key int
	val int
}

type LRUCache struct {
	capacity int
	ll       *list.List
	pos      map[int]*list.Element
}

// Get returs (value, ok)
// If key is exist it becomes MRU (MoveToFron)
func (c *LRUCache) Get(key int) (int, bool) {
	//TODO:
	// 1. Find an element in c.pos.
	elem, ok := c.pos[key]
	// 2. If an element does not exist return 0, false
	if !ok {
		return 0, false
	}
	// 3. If an element exists to MoveToFront
	c.ll.MoveToFront(elem)
	// 4. Take of entry form elem.Value and return entry.value, true
	entry := elem.Value.(*entry)
	return entry.val, true
}

// Put puts/update value
func (c *LRUCache) Put(key, value int) {
	//TODO:
	// If the key is exist update and MoveToFron
	// If the key is new:
	// - if capacity == 0 nothing save
	// - if cache is full: delete LRU (Back)
	// - then return a new key is MRU (PushFront)
	if elem, ok := c.pos[key]; ok {
		c.ll.MoveToFront(elem)
		entry := elem.Value.(*entry)
		entry.val = value
		return
	}
	if c.capacity == 0 {
		return
	}

	if c.ll.Len() == c.capacity {
		lruElem := c.ll.Back()
		lruEntry := lruElem.Value.(*entry)
		delete(c.pos, lruEntry.key)
		c.ll.Remove(lruElem)
	}

	newEntry := &entry{key: key, val: value}
	newElem := c.ll.PushFront(newEntry)
	c.pos[key] = newElem
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var C, Q int

	if _, err := fmt.Fscan(in, &C, &Q); err != nil {
		return
	}

	cache := NewLRUCache(C)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < Q; i++ {
		var cmd string
		fmt.Fscan(in, &cmd)

		if cmd == "PUT" {
			var key, value int
			fmt.Fscan(in, &key, &value)
			cache.Put(key, value)
		} else if cmd == "GET" {
			var key int
			fmt.Fscan(in, &key)
			if val, ok := cache.Get(key); ok {
				fmt.Fprintln(out, val)
			} else {
				fmt.Fprintln(out, -1)
			}
		}
	}
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		pos:      make(map[int]*list.Element),
	}
}
