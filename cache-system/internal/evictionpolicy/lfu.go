package evictionpolicy

import (
	"cache-system/internal/models"
	"container/list"
	"fmt"
)

type LFUCache struct {
	frequencyListMap map[int]*list.List
	elementsFreqMap  map[*list.Element]int
	minFrequency     int
}

func newLFUCache() *LFUCache {
	return &LFUCache{
		frequencyListMap: make(map[int]*list.List),
		elementsFreqMap:  make(map[*list.Element]int),
		minFrequency:     0,
	}
}

func (c *LFUCache) Get(element *list.Element) *list.Element {
	freq, node := c.elementsFreqMap[element], element.Value

	delete(c.elementsFreqMap, element)
	c.frequencyListMap[freq].Remove(element)
	if c.frequencyListMap[freq].Len() == 0 {
		delete(c.frequencyListMap, freq)
		if freq == c.minFrequency {
			c.minFrequency++
		}
	}

	freq++
	if c.frequencyListMap[freq] == nil {
		c.frequencyListMap[freq] = list.New()
	}
	element = c.frequencyListMap[freq].PushFront(node)
	c.elementsFreqMap[element] = freq
	return element
}

func (c *LFUCache) Evict() *list.Element {
	if c.minFrequency == 0 {
		return nil
	}
	list, ok := c.frequencyListMap[c.minFrequency]
	if !ok || list.Len() == 0 {
		return nil
	}

	last := c.frequencyListMap[c.minFrequency].Back()
	if last != nil {
		c.frequencyListMap[c.minFrequency].Remove(last)
		delete(c.elementsFreqMap, last)
	}

	if c.frequencyListMap[c.minFrequency].Len() == 0 {
		delete(c.frequencyListMap, c.minFrequency)
		c.minFrequency++
	}

	return last
}

func (c *LFUCache) Put(node any) *list.Element {
	if c.frequencyListMap[1] == nil {
		c.frequencyListMap[1] = list.New()
	}
	element := c.frequencyListMap[1].PushFront(node)
	c.elementsFreqMap[element] = 1
	c.minFrequency = 1
	return element
}

func (c *LFUCache) Delete(element *list.Element) {
	freq := c.elementsFreqMap[element]
	delete(c.elementsFreqMap, element)
	c.frequencyListMap[freq].Remove(element)

	if c.frequencyListMap[freq].Len() == 0 {
		delete(c.frequencyListMap, freq)
		if c.minFrequency == freq {
			c.minFrequency++
		}
	}
}

// just for debugging
func (c *LFUCache) Print() {
	fmt.Println("LFU Cache State:")
	for freq, l := range c.frequencyListMap {
		fmt.Printf("Frequency %d: ", freq)
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value.(*models.Node[string, string]).GetKey())
		}
		fmt.Println()
	}
}
