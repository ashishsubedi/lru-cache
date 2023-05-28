package main

import "LRUCache/pkg/lrucache"

const SIZE int = 5

func main() {
	lruCache := lrucache.NewCache(SIZE)
	elements := []string{"a", "b", "c", "d", "e", "f", "d"}
	for _, e := range elements {
		lruCache.Check(e)
		lruCache.Display()
	}

}
