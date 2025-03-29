package main

import (
	"cache-system/internal/cache"
	"cache-system/internal/evictionpolicy"
	"fmt"
	"time"
)

func main() {
	RunKeysExpiration()
}

func RunKeysExpiration() {
	lru, _ := evictionpolicy.NewEvictionPolicy(evictionpolicy.LRU)
	cache := cache.NewCache[string, string](4, 2, lru)

	cache.Put("Meowth", "1", 3)
	cache.Put("Ben10", "2", 3)

	fmt.Println("Cache after adding keys:")
	cache.Print()

	time.Sleep(4 * time.Second)

	fmt.Println("Cache after expiration:")
	cache.Print()

	cache.Put("Gwen", "3", 5)
	fmt.Println("Cache after adding a new key:")
	cache.Print()

	time.Sleep(6 * time.Second)
	fmt.Println("Cache after new key expiration:")
	cache.Print()
}

func RunFIFOCache() {
	fifo, _ := evictionpolicy.NewEvictionPolicy(evictionpolicy.FIFO)
	fifoCache := cache.NewCache[string, string](4, 30, fifo)

	fmt.Println("FIFO Cache:")
	fifoCache.Put("Meowth", "1", 300)
	fifoCache.Put("Ben10", "2", 300)
	fifoCache.Put("Gwen", "3", 300)
	fifoCache.Put("Kevin", "4", 300)

	fifoCache.Print()
	fifoCache.Put("Max", "5", 300)
	fifoCache.Put("Ben10", "10", 300)

	fifoCache.Print()
	fifoCache.Put("Pikachu", "6", 300)
	fifoCache.Print()
}

func RunLFUCache() {
	lfu, _ := evictionpolicy.NewEvictionPolicy(evictionpolicy.LFU)
	lfuCache := cache.NewCache[string, string](4, 30, lfu)

	fmt.Println("LFU Cache:")
	lfuCache.Put("Meowth", "1", 300)
	lfuCache.Put("Ben10", "2", 300)
	lfuCache.Put("Gwen", "3", 300)
	lfuCache.Put("Kevin", "4", 300)

	lfuCache.Print()
	lfuCache.Get("Ben10")
	lfuCache.Get("Ben10")
	lfuCache.Put("Ben10", "10", 300)
	lfuCache.Print()
	lfuCache.Get("Gwen")

	lfuCache.Put("Max", "5", 300)
	lfuCache.Print()
	lfuCache.Put("Pikachu", "6", 300)
	lfuCache.Print()
}

func RunLRUCache() {
	lru, _ := evictionpolicy.NewEvictionPolicy(evictionpolicy.LRU)
	lruCache := cache.NewCache[string, string](4, 30, lru)

	fmt.Println("LRU Cache:")
	lruCache.Put("Meowth", "1", 300)
	lruCache.Put("Ben10", "2", 300)
	lruCache.Put("Gwen", "3", 300)
	lruCache.Put("Kevin", "4", 300)

	lruCache.Print()
	lruCache.Put("Max", "5", 300)
	lruCache.Put("Ben10", "10", 300)
	lruCache.Get("Gwen")

	lruCache.Print()
	lruCache.Put("Pikachu", "6", 300)
	lruCache.Print()
}
