package main

import (
	"container/list"
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	sum := md5.Sum([]byte("fasdjklfjasdkljgds"))
	fmt.Printf("%x\n", sum)
	unix := time.Now().Unix()
	fmt.Printf("%x\n", unix)
	fmt.Println(len(`{"x":1234,"y":11,"w":200,"h":200,"prob":0.9999}`))
	str := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	part := str[:1:1]
	fmt.Println(part, len(part), cap(part))
	fmt.Println(str[:0])
	fmt.Println(300001 & 0x01)
	fmt.Println(uint16(1<<15) - 1 + uint16(1<<15))
	i1 := 100
	i2 := i1
	fmt.Println(&i1 == &i2)
	fmt.Println(&i1 == &i1)
	fmt.Println(^uint(0))
	maxInt := int(^uint(0) >> 1)
	fmt.Println(maxInt)
	minInt := ^maxInt
	fmt.Println(uint(minInt))
	fmt.Println(-1 << 63)
	fmt.Println((1 << 63) - 1)

	cache := NewLRUCache(2)
	fmt.Println(cache.Get(1))
	cache.Put(1, 2)
	cache.Put(2, 2)
	cache.Put(3, 2)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
	cache.Put(1, 2)
	fmt.Println(cache.Get(3))
	l := list.New()
	l.Back()
}

type LRUCache struct {
	c   map[int][2]interface{}
	l   *list.List
	len int
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{make(map[int][2]interface{}), list.New(), capacity}
}

func (lru *LRUCache) Get(key int) int {
	if cache, ok := lru.c[key]; ok {
		lru.l.MoveToBack(cache[1].(*list.Element))
		return cache[0].(int)
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if lru.l.Len() >= lru.len {
		elem := lru.l.Front()
		lru.l.Remove(elem)
		delete(lru.c, elem.Value.(int))
	}
	elemPushed := lru.l.PushBack(key)
	lru.c[key] = [2]interface{}{value, elemPushed}
}
