package evictionpolicy

import (
	"cache-system/internal/models"
	"container/list"
	"fmt"
)

type FIFOCache struct {
	list *list.List
}

func newFIFOCache() *FIFOCache {
	return &FIFOCache{
		list: list.New(),
	}
}

func (c *FIFOCache) Get(element *list.Element) *list.Element {
	return element
}

func (c *FIFOCache) Evict() *list.Element {
	front := c.list.Front()
	if front != nil {
		c.list.Remove(front)
	}
	return front
}

func (c *FIFOCache) Put(node any) *list.Element {
	return c.list.PushBack(node)
}

func (c *FIFOCache) Delete(element *list.Element) {
	if element != nil {
		c.list.Remove(element)
	}
}

// just for debugging
func (c *FIFOCache) Print() {
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(*models.Node[string, string]).GetKey())
	}
	fmt.Println()
}
