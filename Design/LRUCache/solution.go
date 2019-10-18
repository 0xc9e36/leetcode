/**
 * @Author: wei.tan
 * @Description:
 * @File:  Design
 * @Version: 1.0.0
 * @Date: 2019-10-18 23:14
 */

package LRUCache


type Node struct {
	key, value int
	prev, next *Node
}


type LRUCache struct {
	capacity   int
	cacheMap   map[int]*Node
	head, tail *Node
}


func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.tail = &Node{}
	lru.head = &Node{}
	lru.capacity = capacity
	lru.cacheMap = make(map[int]*Node)
	lru.head.next = lru.tail
	lru.head.prev = nil
	lru.tail.next = nil
	lru.tail.prev = lru.head
	return lru
}

func (this LRUCache) delete(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this LRUCache) attach(node *Node) {
	node.next = this.head.next
	node.next.prev = node
	this.head.next = node
	node.prev = this.head
}


func (this LRUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}

	this.delete(node)
	this.attach(node)
	return node.value
}


func (this LRUCache) Put(key int, value int)  {

	oldNode, ok := this.cacheMap[key]

	//如果存在 key
	if ok {
		this.delete(oldNode)
		this.attach(oldNode)
		oldNode.value = value
	} else {
		var node *Node
		//没有剩余空间, 删除最近最少使用节点
		if len(this.cacheMap) >= this.capacity {
			node = this.tail.prev
			this.delete(node)
			delete(this.cacheMap, node.key)
		} else {
			node = new(Node)
		}

		node.key = key
		node.value = value

		this.cacheMap[key] = node

		//放到头结点
		this.attach(node)
	}
}
