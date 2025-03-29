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
	mtx            sync.Mutex
	capacity       int
	expireInterval int
	storage        map[T]*list.Element
	evictionpolicy evictionpolicy.EvictionPolicy
}

func NewCache[T comparable, U any](capacity, expireInterval int, evictionpolicy evictionpolicy.EvictionPolicy) *Cache[T, U] {
	cache := &Cache[T, U]{
		capacity:       capacity,
		expireInterval: expireInterval,
		storage:        make(map[T]*list.Element),
		evictionpolicy: evictionpolicy,
	}
	go cache.expire()
	return cache
}

func (c *Cache[T, U]) Get(key T) (U, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if element, ok := c.storage[key]; ok {
		element = c.evictionpolicy.Get(element)
		c.storage[key] = element
		if node, ok := element.Value.(*models.Node[T, U]); ok {
			if !node.IsExpire() {
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
		element = c.evictionpolicy.Get(element)
		c.storage[key] = element
	} else {
		if c.capacity <= len(c.storage) {
			del := c.evictionpolicy.Evict()
			if del != nil {
				delete(c.storage, del.Value.(*models.Node[T, U]).GetKey())
			}
		}
		c.storage[key] = c.evictionpolicy.Put(models.NewNode(key, value, ttl))
	}
}

func (c *Cache[T, U]) Delete(key T) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.delete(key)
}

func (c *Cache[T, U]) expire() {
	ticker := time.NewTicker(time.Duration(c.expireInterval) * time.Second)
	fmt.Printf("Cache expire interval started, interval: %v\n", c.expireInterval)

	for range ticker.C {
		c.mtx.Lock()
		for key, value := range c.storage {
			if value.Value.(*models.Node[T, U]).IsExpire() {
				c.delete(key)
			}
		}
		c.mtx.Unlock()
	}
}

func (c *Cache[T, U]) delete(key T) {
	if element, ok := c.storage[key]; ok {
		c.evictionpolicy.Delete(element)
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
	c.evictionpolicy.Print()
}
