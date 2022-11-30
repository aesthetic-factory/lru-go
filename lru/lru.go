package lru

import (
	"errors"
	"fmt"
)

type Key interface {
	int | int64 | int32 | string
}

type DLinkListNode[K Key, V any] struct {
	prev  *DLinkListNode[K, V]
	next  *DLinkListNode[K, V]
	key   K
	value V
}

type LRU[K Key, V any] struct {
	head    *DLinkListNode[K, V]
	tail    *DLinkListNode[K, V]
	hmap    map[K]*DLinkListNode[K, V]
	maxSize int
	size    int
}

// set node as head
func (lru *LRU[K, V]) pushFront(node *DLinkListNode[K, V]) {
	if lru.head == nil || lru.tail == nil {
		lru.head = node
		lru.tail = node
	} else {
		oldHead := lru.head
		lru.head = node
		lru.head.next = oldHead
		oldHead.prev = lru.head
	}
}

// remove node from list, connect n-1 and n+1 node
func (lru *LRU[K, V]) detach(node *DLinkListNode[K, V]) {

	if lru.tail == node {
		lru.tail = node.prev
	}
	nodePrev := node.prev
	nodeNext := node.next

	if nodePrev != nil {
		nodePrev.next = nodeNext
	}
	if nodeNext != nil {
		nodeNext.prev = nodePrev
	}
}

// init hash map and set size
func (lru *LRU[K, V]) Init(maxSize int) {
	lru.hmap = make(map[K]*DLinkListNode[K, V])
	lru.maxSize = maxSize
	lru.size = 0
}

// insert new item, or replace existing item with same key
func (lru *LRU[K, V]) Insert(key K, value V) {
	n, exist := lru.hmap[key]

	if exist {
		n.value = value
		lru.detach(n)
		lru.pushFront(n)
	} else {
		node := new(DLinkListNode[K, V])
		node.key = key
		node.value = value
		lru.hmap[key] = node
		lru.pushFront(node)

		lru.size++
		if lru.size > lru.maxSize {
			lru.Remove(lru.tail.key)
		}
	}
}

// get item by key
func (lru *LRU[K, V]) Get(key K) (V, error) {
	n, exist := lru.hmap[key]
	if exist {
		lru.detach(n)
		lru.pushFront(n)
		return n.value, nil
	}
	var dum V
	return dum, errors.New("not found")
}

// remove item by key
func (lru *LRU[K, V]) Remove(key K) {
	n, exist := lru.hmap[key]
	if exist {
		lru.detach(n)
		delete(lru.hmap, key)
		lru.size--
	}
}

// display linklist info for debug
func (lru *LRU[K, V]) Show() {
	fmt.Println("Display linklist")
	n := lru.head
	for n != nil {
		fmt.Println(n.key, " ", n.value)
		if n == lru.tail {
			fmt.Println("Reached tail")
		}
		n = n.next
	}
	fmt.Println("") // print new line
}
