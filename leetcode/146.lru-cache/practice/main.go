package practice

// 缺少 DLinkedNode 的构造函数
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
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
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	// fixme: 初始化的时候要将 head 和 tail 相连
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// 使用卫述语句更好
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		value := node.value
		this.moveToHead(key, value)
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; ok {
		node := this.cache[key]
		// fixme: 这里只能更新 value 一个值。 通过 node.value = value
		node.value = value

		this.moveToHead(key, value)
		return // fixme: 忘记 return
	}
	this.addToHead(key, value)
}

// 入参为 node 更好
func (this *LRUCache) addToHead(key int, value int) {
	node := &DLinkedNode{
		key:   key,
		value: value,
	}

	this.cache[key] = node

	//TODO: addToHead 顺序梳理
	node.next = this.head.next
	this.head.next.prev = node

	this.head.next = node
	node.prev = this.head

	// fixme: 没有更新 size
	this.size++
	if this.size > this.capacity {
		this.removeTail()
	}
}

// 入参为 node 更好
func (this *LRUCache) moveToHead(key int, value int) {
	node := this.cache[key]
	this.removeNode(node)
	this.addToHead(key, value)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev

	// fixme: 忘记删掉 map 的 key 了
	delete(this.cache, node.key)
	this.size--
}

func (this *LRUCache) removeTail() {
	node := this.tail.prev
	this.removeNode(node)
}
