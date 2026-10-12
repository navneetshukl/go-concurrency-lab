package main

type node struct {
	val  int
	prev *node
	next *node
	key  int
}

func newNode(key, val int) *node {
	return &node{
		val:  val,
		prev: nil,
		next: nil,
		key:  key,
	}
}

type LRUCache struct {
	mp   map[int]*node
	head *node
	tail *node
	cap  int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cap:  capacity,
		head: newNode(-1, -1),
		tail: newNode(-1, -1),
		mp:   make(map[int]*node),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l

}

func (l *LRUCache) AddNode(n *node) {
	l.head.next.prev = n
	n.next = l.head.next
	l.head.next = n
	n.prev = l.head

}

func (l *LRUCache) DeleteNode(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.mp[key]; !ok {
		return -1
	}
	vall, _ := this.mp[key]
	// delete this node and add at front
	this.DeleteNode(vall)
	this.AddNode(vall)
	return vall.val

}

func (this *LRUCache) Put(key int, value int) {
	//if key is present
	vall, ok := this.mp[key]
	if ok {
		this.DeleteNode(vall)
		vall.val = value
		this.AddNode(vall)
		delete(this.mp, key)
		this.mp[key] = this.head.next

	} else {
		if len(this.mp) == this.cap {
			delete(this.mp, this.tail.prev.key)
			this.DeleteNode(this.tail.prev)
		}
		no := newNode(key, value)
		this.AddNode(no)
		this.mp[key] = no
	}
}

func main() {
	// Test case:
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	// Expected: [null, null, null, 1, null, -1, null, -1, 3, 4]

	cache := Constructor(2)
	
	cache.Put(1, 1)
	cache.Put(2, 2)
	println(cache.Get(1)) // 1
	cache.Put(3, 3)
	println(cache.Get(2)) // -1
	cache.Put(4, 4)
	println(cache.Get(1)) // -1
	println(cache.Get(3)) // 3
	println(cache.Get(4)) // 4
}
