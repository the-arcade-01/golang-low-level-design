package cache

import (
	"cache-system/internal/evictionpolicy"
	"cache-system/internal/models"
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Cache[T comparable, U any] struct {
	capacity       int
	storage        map[T]*list.Element
	evictionPolicy evictionpolicy.EvictionPolicy
	mtx            sync.RWMutex
}

func NewCache[T comparable, U any](capacity int, evictionPolicy evictionpolicy.EvictionPolicy) *Cache[T, U] {
	return &Cache[T, U]{
		capacity:       capacity,
		storage:        make(map[T]*list.Element),
		evictionPolicy: evictionPolicy,
	}
}

func (c *Cache[T, U]) Get(key T) (U, error) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	if element, ok := c.storage[key]; ok {
		element = c.evictionPolicy.Get(element)
		c.storage[key] = element
		if node, ok := element.Value.(*models.Node[T, U]); ok {
			if node.GetTTL().After(time.Now()) {
				return node.GetValue(), nil
			}
		}
	}
	var zeroValue U
	return zeroValue, fmt.Errorf("key not found")
}

func (c *Cache[T, U]) Put(key T, value U, ttl int) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.capacity == 0 {
		return
	}

	if element, ok := c.storage[key]; ok {
		element.Value.(*models.Node[T, U]).UpdateValue(value).UpdateTTL(ttl)
		c.evictionPolicy.Get(element)
	} else {
		if c.capacity <= len(c.storage) {
			last := c.evictionPolicy.Evict()
			if last != nil {
				delete(c.storage, last.Value.(*models.Node[T, U]).GetKey())
			}
		}
		c.storage[key] = c.evictionPolicy.Put(models.NewNode(key, value, ttl))
	}
}

func (c *Cache[T, U]) Delete(key T) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if element, ok := c.storage[key]; ok {
		c.evictionPolicy.Delete(element)
		delete(c.storage, key)
	}
}

// just for debugging
func (c *Cache[T, U]) Print() {
	fmt.Println("Storage")
	for key := range c.storage {
		fmt.Printf("%v ", key)
	}
	fmt.Println("\nList")
	c.evictionPolicy.Print()
}
