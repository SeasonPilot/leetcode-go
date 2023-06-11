package practice4

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; !ok {
		return -1
	} else {
		l.moveToHead(node)
		return node.value
	}
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.moveToHead(node)
	} else {
		newNode := &DLinkedNode{
			key:   key,
			value: value,
			//prev:  &DLinkedNode{},//fixme: 不需要初始化，l.addToHead(newNode) 不会访问 newNode.prev，只会给 newNode.prev 赋值
			//next:  &DLinkedNode{},
		}

		l.addToHead(newNode)
		l.cache[key] = newNode
		l.size++

		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.cache, removed.key) //fixme: removed.key
			l.size--
		}
	}
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node) //fixme:顺序不能颠倒，一定要先 removeNode 再 addToHead
	l.addToHead(node)
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.next = l.head.next // 将新节点的next指向head的下一个节点
	l.head.next.prev = node // 将head的下一个节点的prev指向新节点

	node.prev = l.head // 将新节点的prev指向head
	l.head.next = node // 将head的next指向新节点
}

func (l *LRUCache) removeTail() *DLinkedNode {
	removed := l.tail.prev //
	l.removeNode(removed)
	return removed
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
