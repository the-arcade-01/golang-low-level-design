package main

import (
	"cache-system/internal/cache"
	"cache-system/internal/evictionpolicy"
	"fmt"
)

func main() {
	lru, _ := evictionpolicy.NewEvictionPolicy("LRU")
	lruCache := cache.NewCache[string, string](4, lru)

	fmt.Println("LRU Cache:")
	lruCache.Put("Meowth", "1", 300)
	lruCache.Put("Ben10", "2", 300)
	lruCache.Put("Gwen", "3", 300)
	lruCache.Put("Kevin", "4", 300)

	lruCache.Print()
	lruCache.Put("Max", "5", 300)
	lruCache.Get("Ben10")

	lruCache.Print()
	lruCache.Put("Pikachu", "6", 300)
	lruCache.Print()

	fifo, _ := evictionpolicy.NewEvictionPolicy("FIFO")
	fifoCache := cache.NewCache[string, string](4, fifo)

	fmt.Println("FIFO Cache:")
	fifoCache.Put("Meowth", "1", 300)
	fifoCache.Put("Ben10", "2", 300)
	fifoCache.Put("Gwen", "3", 300)
	fifoCache.Put("Kevin", "4", 300)

	fifoCache.Print()
	fifoCache.Put("Max", "5", 300)
	fifoCache.Get("Ben10")

	fifoCache.Print()
	fifoCache.Put("Pikachu", "6", 300)
	fifoCache.Print()

	lfu, _ := evictionpolicy.NewEvictionPolicy("LFU")
	lfuCache := cache.NewCache[string, string](4, lfu)

	fmt.Println("LFU Cache:")
	lfuCache.Put("Meowth", "1", 300)
	lfuCache.Put("Ben10", "2", 300)
	lfuCache.Put("Gwen", "3", 300)
	lfuCache.Put("Kevin", "4", 300)

	lfuCache.Print()
	lfuCache.Get("Ben10")
	lfuCache.Get("Ben10")
	lfuCache.Print()
	lfuCache.Get("Gwen")

	lfuCache.Put("Max", "5", 300)
	lfuCache.Print()
	lfuCache.Put("Pikachu", "6", 300)
	lfuCache.Print()
}
