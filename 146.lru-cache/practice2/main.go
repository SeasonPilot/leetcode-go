package practice2

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func NewDLinkedNode(key int, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
		prev:  new(DLinkedNode),
		next:  new(DLinkedNode),
	}
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkedNode
	head, tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     NewDLinkedNode(0, 0),
		tail:     NewDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.moveToHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		node = NewDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	this.head.next.prev = node

	this.head.next = node
	node.prev = this.head

	this.size++
	if this.size > this.capacity {
		this.removeTail()
	}
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() {
	node := this.tail.prev
	this.removeNode(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	this.size--

	delete(this.cache, node.key)
}
