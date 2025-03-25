package models

import "time"

type Node[T comparable, U any] struct {
	key   T
	value U
	ttl   time.Time
}

func NewNode[T comparable, U any](key T, value U, ttl int) *Node[T, U] {
	return &Node[T, U]{
		key:   key,
		value: value,
		ttl:   time.Now().Add(time.Duration(ttl) * time.Second),
	}
}

func (n *Node[T, U]) GetKey() T {
	return n.key
}

func (n *Node[T, U]) GetTTL() time.Time {
	return n.ttl
}

func (n *Node[T, U]) GetValue() U {
	return n.value
}

func (n *Node[T, U]) UpdateValue(value U) *Node[T, U] {
	n.value = value
	return n
}

func (n *Node[T, U]) UpdateTTL(ttl int) *Node[T, U] {
	n.ttl = time.Now().Add(time.Duration(ttl) * time.Second)
	return n
}
