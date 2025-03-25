package evictionpolicy

import (
	"cache-system/internal/models"
	"container/list"
	"fmt"
)

type LRU struct {
	list *list.List
}

func newLRU() *LRU {
	return &LRU{
		list: list.New(),
	}
}

func (c *LRU) Evict() *list.Element {
	last := c.list.Back()
	if last != nil {
		c.list.Remove(last)
	}
	return last
}

func (c *LRU) Get(element *list.Element) *list.Element {
	c.list.MoveToFront(element)
	return element
}

func (c *LRU) Put(node any) *list.Element {
	return c.list.PushFront(node)
}

func (c *LRU) Delete(element *list.Element) {
	c.list.Remove(element)
}

// just for debugging
func (c *LRU) Print() {
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(*models.Node[string, string]).GetKey())
	}
	fmt.Println()
}
