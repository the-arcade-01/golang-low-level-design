package evictionpolicy

import (
	"container/list"
	"fmt"
)

type PolicyType string

const (
	LRU  PolicyType = "LRU"
	FIFO PolicyType = "FIFO"
	LFU  PolicyType = "LFU"
)

type EvictionPolicy interface {
	Get(element *list.Element) *list.Element
	Evict() *list.Element
	Put(node any) *list.Element
	Delete(element *list.Element)
	Print()
}

func NewEvictionPolicy(policy PolicyType) (EvictionPolicy, error) {
	switch policy {
	case LRU:
		return newLRUCache(), nil
	case FIFO:
		return newFIFOCache(), nil
	case LFU:
		return newLFUCache(), nil
	}

	return nil, fmt.Errorf("policy not found")
}
