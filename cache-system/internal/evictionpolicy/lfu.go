package evictionpolicy

import (
	"cache-system/internal/models"
	"container/list"
	"fmt"
)

type LFU struct {
	frequencyListMap map[int]*list.List
	elementsFreq     map[*list.Element]int
	minFreq          int
}

func newLFU() *LFU {
	return &LFU{
		frequencyListMap: make(map[int]*list.List),
		elementsFreq:     make(map[*list.Element]int),
		minFreq:          0,
	}
}

func (c *LFU) Evict() *list.Element {
	if c.minFreq == 0 {
		return nil
	}
	newList, ok := c.frequencyListMap[c.minFreq]
	if !ok || newList.Len() == 0 {
		return nil
	}
	element := newList.Back()
	newList.Remove(element)
	delete(c.elementsFreq, element)

	if newList.Len() == 0 {
		delete(c.frequencyListMap, c.minFreq)
		c.minFreq++
	}
	return element
}

func (c *LFU) Get(element *list.Element) *list.Element {
	node, freq := element.Value, c.elementsFreq[element]
	delete(c.elementsFreq, element)
	c.frequencyListMap[freq].Remove(element)

	if c.frequencyListMap[freq].Len() == 0 {
		delete(c.frequencyListMap, freq)
		if c.minFreq == freq {
			c.minFreq++
		}
	}

	freq++
	if c.frequencyListMap[freq] == nil {
		c.frequencyListMap[freq] = list.New()
	}
	newElement := c.frequencyListMap[freq].PushFront(node)
	c.elementsFreq[newElement] = freq
	return newElement
}

func (c *LFU) Put(node any) *list.Element {
	if c.frequencyListMap[1] == nil {
		c.frequencyListMap[1] = list.New()
	}
	element := c.frequencyListMap[1].PushFront(node)
	c.elementsFreq[element] = 1
	c.minFreq = 1
	return element
}

func (c *LFU) Delete(element *list.Element) {
	freq := c.elementsFreq[element]
	delete(c.elementsFreq, element)
	c.frequencyListMap[freq].Remove(element)

	if c.frequencyListMap[freq].Len() == 0 {
		delete(c.frequencyListMap, freq)
		if c.minFreq == freq {
			c.minFreq++
		}
	}
}

// just for debugging
func (c *LFU) Print() {
	fmt.Println("LFU Cache State:")
	for freq, l := range c.frequencyListMap {
		fmt.Printf("Frequency %d: ", freq)
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value.(*models.Node[string, string]).GetKey())
		}
		fmt.Println()
	}
}
