package main

import "fmt"

type node struct {
	value     interface{}
	pre, next *node
}

type list struct {
	len  int
	head *node
	tail *node
}

func newNode(v interface{}) *node {
	return &node{value: v, pre: nil, next: nil}
}

func (l *list) Init() {
	l.len = 0
	l.head, l.tail = nil, nil
}

func (l *list) LPush(v interface{}) {
	node := newNode(v)
	if l.len == 0 {
		l.tail = node
	} else {
		node.next = l.head
		l.head.pre = node
	}
	l.head = node
	l.len++
	return
}

func (l *list) RPush(v interface{}) {
	node := newNode(v)
	if l.len == 0 {
		l.head = node
	} else {
		node.pre = l.tail
		l.tail.next = node
	}
	l.tail = node
	l.len++
	return
}

func (l *list) LPop() *node {
	if l.len == 0 {
		return nil
	}
	n := l.head
	l.head = l.head.next
	l.len--
	return n
}

func (l *list) RPop() *node {
	if l.len == 0 {
		return nil
	}
	n := l.tail
	l.tail = l.tail.pre
	l.tail.next = nil

	l.len--
	return n
}

func (l *list) Reset() {
	l.len = 0
	l.head, l.tail = nil, nil
}

func (l *list) Merge(r *list) {

	if l.len == 0 {
		l = r
		return
	}
	l.tail.next = r.head
	l.len += r.len
	return
}

//TODO

func (l *list) Reverse() {
	if l.len <= 1 {
		return
	}

	root := l.head
	l.head = l.tail
	l.tail = root
	var pre *node
	var next *node
	for root.next != nil {
		next = root.next
		root.next = pre
		root.pre = next
		pre = root
		root = next
	}
	root.next = pre
	return
}

func main() {
	l := &list{}
	l.Init()
	for i := 0; i < 10; i++ {
		l.LPush(i)
	}
	j := 0
	for cur := l.head; cur != nil; cur = cur.next {
		j++
		fmt.Printf("Index %d node: value => %+v\n", j, cur.value)
	}
	n := l.RPop()
	fmt.Printf("RPop value => %+v\n", n.value)
	j = 0
	for cur := l.head; cur != nil; cur = cur.next {
		j++
		fmt.Printf("Index %d node: value => %+v\n", j, cur.value)
	}
	for i := 11; i < 20; i++ {
		l.RPush(i)
	}
	j = 0
	for cur := l.head; cur != nil; cur = cur.next {
		j++
		fmt.Printf("Index %d node: value => %+v\n", j, cur.value)
	}

	n = l.LPop()
	fmt.Printf("LPop => %+v, len => %d\n", n.value, l.len)

	l.Reverse()
	j = 0
	for cur := l.head; cur != nil; cur = cur.next {
		j++
		fmt.Printf("Index(reverse) %d node: value => %+v\n", j, cur.value)
	}

}
