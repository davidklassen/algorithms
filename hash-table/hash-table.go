package main

import (
	"fmt"
)

type nodeList struct {
	key   string
	value interface{}
	next  *nodeList
}

func (nl *nodeList) insert(key string, value interface{}) *nodeList {
	return &nodeList{
		key:   key,
		value: value,
		next:  nl,
	}
}

func (nl *nodeList) find(key string) interface{} {
	curr := nl
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.next
	}
	return nil
}

type HashTable struct {
	arr []*nodeList
}

func NewHashTable(cap int) *HashTable {
	return &HashTable{arr: make([]*nodeList, cap)}
}

func (ht *HashTable) Get(key string) interface{} {
	bucket := ht.arr[ht.hash(key)%len(ht.arr)]
	if bucket == nil {
		return nil
	}
	return bucket.find(key)
}

func (ht *HashTable) Set(key string, value interface{}) {
	bucket := &ht.arr[ht.hash(key)%len(ht.arr)]
	if *bucket == nil {
		*bucket = &nodeList{key: key, value: value}
	} else {
		*bucket = (*bucket).insert(key, value)
	}
}

func (ht *HashTable) hash(key string) int {
	return 1
}

func main() {
	fmt.Println("Hello, World!")
	m := NewHashTable(100)
	m.Set("foo", 123)
	m.Set("bar", 555)
	fmt.Println(m.Get("foo"))
	fmt.Println(m.Get("bar"))
	fmt.Println(m.Get("buz"))
}
