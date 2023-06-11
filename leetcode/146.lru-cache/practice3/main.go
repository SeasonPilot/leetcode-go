package main

//只需要对外暴露 LRUCache、DLinkedNode结构体，以及LRUCache的 GET 和 PUT 方法、构造函数即可
type LRUCache struct {
	cache          map[int]*DLinkedNode
	size, capacity int // LRUCache 是有容量和大小的
	head, tail     *DLinkedNode
}

// 字段都是小写,链表是DLinkedNode
type DLinkedNode struct {
	key, value int //key 是干啥用的？  删除 cache 中的 key 需要从移除的 node 中获取 key
	prev, next *DLinkedNode
}

func initDLinkedNode(key, val int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: val,
		//prev:  new(DLinkedNode), //不需要初始化
		//next:  new(DLinkedNode), //
	}
}

// fixme: 需要传入 capacity，且只传 capacity 即可
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: make(map[int]*DLinkedNode, capacity), //
		//cache:    map[int]*DLinkedNode{},
		capacity: capacity,
		head:     initDLinkedNode(0, 0), //使用new(DLinkedNode)不行吗？
		tail:     initDLinkedNode(0, 0),
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

func (l *LRUCache) addToHead(node *DLinkedNode) {
	next := l.head.next // 这边不需要存储
	node.next = next
	next.prev = node

	l.head.next = node
	node.prev = l.head
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) Put(key int, value int) {
	//fixme: put 要检查 key 对应的 node 是否存在
	if node, ok := l.cache[key]; ok {
		node.value = value //fixme: node.key = value 是错误的
		l.moveToHead(node)
	} else {
		newNode := initDLinkedNode(key, value) // fixme: key 在 DLinkedNode、LRUCache 中都存在，所以 put 的时候都需要设置值
		l.cache[key] = newNode
		l.addToHead(newNode)
		l.size++ //fixme:size++，增删 cache 中的 node 的时候，需要操作 size

		if l.size > l.capacity {
			delNode := l.removeTail()    //fixme: removeTail() 需要返回被 remove 的 node，不然不知道要删除 map 的 key是什么
			delete(l.cache, delNode.key) //fixme:remove 双链表中的 node，但是 map 中还有该node，所以需要删除 map 中的 key
			l.size--
		}
	}
}

func (l *LRUCache) removeTail() *DLinkedNode { //fixme:要返回被 remove 的 node。removeTail()方法需要返回被删除的节点。理由:删除尾节点后,需要从map中删除相应的key,而key只能从被删除的节点中获取。建议:修改removeTail()方法,返回被删除的节点。
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	// 不用管 remove 的 node，直接把 node 前后的节点连接起来即可
	node.prev.next = node.next
	node.next.prev = node.prev
}

//注意： 增删双链表中的节点时，需要记得同时操作 cache 的 key 和 size
