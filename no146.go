package leetcode

// 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

// 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
// 写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

//

// 进阶:

// 你是否可以在 O(1) 时间复杂度内完成这两种操作？

//

// 示例:

// LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );

// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // 返回  1
// cache.put(3, 3);    // 该操作会使得关键字 2 作废
// cache.get(2);       // 返回 -1 (未找到)
// cache.put(4, 4);    // 该操作会使得关键字 1 作废
// cache.get(1);       // 返回 -1 (未找到)
// cache.get(3);       // 返回  3
// cache.get(4);       // 返回  4

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/lru-cache
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
type LRUCache struct {
	size     int
	capacity int
	index    map[int]*CacheNode
	head     *CacheNode
	tail     *CacheNode
}

type CacheNode struct {
	key   int
	value int
	prev  *CacheNode
	next  *CacheNode
}

func Constructor(capacity int) LRUCache {
	head := &CacheNode{}
	tail := &CacheNode{
		prev: head,
	}
	head.next = tail
	c := LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		index:    make(map[int]*CacheNode),
	}
	return c
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.index[key]
	if !ok {
		return -1
	}
	this.move2Head(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.index[key]
	if !ok {
		//add
		node := &CacheNode{
			key:   key,
			value: value,
		}
		this.index[key] = node
		this.add2Head(node)
		this.size++
		if this.size > this.capacity {
			this.removeTail()
		}
	} else {
		node.value = value
		this.move2Head(node)
	}

}

func (this *LRUCache) move2Head(node *CacheNode) {
	this.removeNode(node)
	this.add2Head(node)
}

func (this *LRUCache) add2Head(node *CacheNode) {
	node.next = this.head.next
	node.prev = this.head
	node.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *CacheNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

func (this *LRUCache) removeTail() {
	last := this.tail.prev
	this.removeNode(last)
	delete(this.index, last.key)
	this.size--
}

// type LRUCache struct {
// 	size       int
// 	capacity   int
// 	cache      map[int]*DLinkedNode
// 	head, tail *DLinkedNode
// }

// type DLinkedNode struct {
// 	key   int
// 	value int
// 	prev  *DLinkedNode
// 	next  *DLinkedNode
// }

// func Constructor(capacity int) LRUCache {
// 	cache := LRUCache{
// 		size:     0,
// 		capacity: capacity,
// 		cache:    make(map[int]*DLinkedNode),
// 	}

// 	head := &DLinkedNode{}
// 	tail := &DLinkedNode{}
// 	tail.prev = head
// 	head.next = tail
// 	cache.head = head
// 	cache.tail = tail

// 	return cache
// }

// func (this *LRUCache) Get(key int) int {
// 	node, ok := this.cache[key]
// 	if ok {
// 		this.move2Head(node)
// 		return node.value
// 	}

// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	node, ok := this.cache[key]
// 	if ok {
// 		node.value = value
// 		this.move2Head(node)
// 		return
// 	}

// 	newNode := &DLinkedNode{
// 		key:   key,
// 		value: value,
// 	}

// 	this.cache[key] = newNode
// 	this.add2Head(newNode)
// 	this.size++
// 	if this.size > this.capacity {
// 		this.removeTail()

// 	}

// }

// func (this *LRUCache) removeTail() {
// 	if this.tail.prev == this.head {
// 		return
// 	}
// 	delete(this.cache, this.tail.prev.key)
// 	this.removeNode(this.tail.prev)

// }

// func (this *LRUCache) move2Head(node *DLinkedNode) {
// 	if nil == node {
// 		return
// 	}

// 	this.removeNode(node)
// 	this.add2Head(node)
// }

// func (this *LRUCache) removeNode(node *DLinkedNode) {
// 	if nil == node {
// 		return
// 	}

// 	node.prev.next = node.next
// 	node.next.prev = node.prev
// }

// func (this *LRUCache) add2Head(node *DLinkedNode) {
// 	if nil == node {
// 		return
// 	}

// 	node.next = this.head.next
// 	node.next.prev = node
// 	this.head.next = node
// 	node.prev = this.head
// }

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
