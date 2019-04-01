/*
   map and double linked list
*/

package main

import "sync"

type Node struct {
	val       interface{}
	pre, next *Node
}

type LRUCache struct {
	sync.RWMutex
	hashmap    map[interface{}]*Node
	capacity   int
	len        int
	head, tail *Node
}

type Cacher interface {
	Init()
	PushFront(n *Node)
	RemoveTail()
}

const (
	MAX_LIST_LEN = 100
)

func NewLRUCache() Cacher {
	return &LRUCache{}
}

func (l *LRUCache) Init() {
	l.hashmap = make(map[interface{}]*Node)
	l.capacity = MAX_LIST_LEN
	l.len = 0
	l.head = nil
	l.tail = nil
}

//TODO:
func (l *LRUCache) PushFront(n *Node) {
	l.Lock()
	defer l.Unlock()
	if l.len == 0 {
		l.tail = n
	}
	if l.len == l.capacity {
		l.RemoveTail()
	}
	n.next = l.head
	l.head = n
	l.len++
	return
}

func (l *LRUCache) RemoveTail() {
	if l.len > 0 {
		l.tail = l.tail.pre
		l.len--
		if l.len == 0 {
			l.head = nil
			l.tail = nil
		}
	}
	return
}

func main() {

}
