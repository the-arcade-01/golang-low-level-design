package evictionpolicy

import (
	"cache-system/internal/models"
	"container/list"
	"fmt"
)

type FIFO struct {
	list *list.List
}

func newFIFO() *FIFO {
	return &FIFO{
		list: list.New(),
	}
}

func (c *FIFO) Evict() *list.Element {
	first := c.list.Front()
	if first != nil {
		c.list.Remove(first)
	}
	return first
}

func (c *FIFO) Get(element *list.Element) *list.Element {
	return element
}

func (c *FIFO) Put(node any) *list.Element {
	return c.list.PushBack(node)
}

func (c *FIFO) Delete(element *list.Element) {
	c.list.Remove(element)
}

// just for debugging
func (c *FIFO) Print() {
	for e := c.list.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value.(*models.Node[string, string]).GetKey())
	}
	fmt.Println()
}
