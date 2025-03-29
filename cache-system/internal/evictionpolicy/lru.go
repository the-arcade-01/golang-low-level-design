package evictionpolicy

import (
	"cache-system/internal/models"
	"container/list"
	"fmt"
)

type LRUCache struct {
	list *list.List
}

func newLRUCache() *LRUCache {
	return &LRUCache{
		list: list.New(),
	}
}

func (c *LRUCache) Get(element *list.Element) *list.Element {
	c.list.MoveToFront(element)
	return element
}

func (c *LRUCache) Evict() *list.Element {
	last := c.list.Back()
	if last != nil {
		c.list.Remove(last)
	}
	return last
}

func (c *LRUCache) Put(node any) *list.Element {
	return c.list.PushFront(node)
}

func (c *LRUCache) Delete(element *list.Element) {
	if element != nil {
		c.list.Remove(element)
	}
}

// just for debugging
func (c *LRUCache) Print() {
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(*models.Node[string, string]).GetKey())
	}
	fmt.Println()
}
