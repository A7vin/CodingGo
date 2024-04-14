package main

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if element, ok := this.cache[key]; ok {
		this.list.MoveToFront(element)
		return element.Value.(pair).value
	}
	return -1
}

func (this *LRUCache) Put(key, value int) {
	if element, ok := this.cache[key]; ok {
		this.list.MoveToFront(element)
		element.Value = pair{key, value}
		return
	}

	if this.list.Len() == this.capacity {
		oldest := this.list.Back()
		delete(this.cache, oldest.Value.(pair).key)
		this.list.Remove(oldest)
	}
	
	element := this.list.PushFront(pair{key, value})
	this.cache[key] = element
}

func main() {
	// 示例代码，创建一个容量为2的LRU缓存
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	println(cache.Get(1)) // 返回 1
	cache.Put(3, 3)       // 使得密钥 2 作废
	println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)       // 使得密钥 1 作废
	println(cache.Get(1)) // 返回 -1 (未找到)
	println(cache.Get(3)) // 返回 3
	println(cache.Get(4)) // 返回 4
}
