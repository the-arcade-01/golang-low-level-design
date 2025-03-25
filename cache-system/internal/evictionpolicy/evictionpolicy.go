package evictionpolicy

import (
	"container/list"
	"fmt"
)

type EvictionPolicy interface {
	Evict() *list.Element
	Get(element *list.Element) *list.Element
	Put(node any) *list.Element
	Delete(element *list.Element)
	Print()
}

func NewEvictionPolicy(policy string) (EvictionPolicy, error) {
	switch policy {
	case "LRU":
		return newLRU(), nil
	case "FIFO":
		return newFIFO(), nil
	case "LFU":
		return newLFU(), nil
	}

	return nil, fmt.Errorf("policy not found")
}
