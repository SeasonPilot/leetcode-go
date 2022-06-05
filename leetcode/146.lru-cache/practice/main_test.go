package practice

import (
	"fmt"
	"testing"
)

func TestLRUCache_Put(t *testing.T) {
	type fields struct {
		size     int
		capacity int
		cache    map[int]*DLinkedNode
		head     *DLinkedNode
		tail     *DLinkedNode
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "put1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ["LRUCache","put","put","get","put","put","get"]
			// [[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
			lru := Constructor(2)
			lru.Put(2, 1)
			lru.Put(2, 2)
			fmt.Println(lru.Get(2))
			lru.Put(1, 1)
			lru.Put(4, 1)
			fmt.Println(lru.Get(2))
		})
	}
}

// nil的变量不可以访问。  nil的变量可以赋值,map类型 除外
// https://fcxnw9u8pt.feishu.cn/docs/doccnhPXOJ7IsmhYmpCekas9fhc#tfppqE
// https://fcxnw9u8pt.feishu.cn/docs/doccnEhN9liZkuFrwLI8grIjPAc#ZlY8pP
func TestDLinkedNode(t *testing.T) {
	// 1.
	node := &DLinkedNode{ // &practice.DLinkedNode{key:1, value:232, prev:(*practice.DLinkedNode)(nil), next:(*practice.DLinkedNode)(nil)}
		key:   1,
		value: 232,
	}

	// 2.nil
	var node *DLinkedNode // panic: runtime error: invalid memory address or nil pointer dereference

	// 3.
	node := new(DLinkedNode) // &practice.DLinkedNode{key:0, value:0, prev:(*practice.DLinkedNode)(nil), next:(*practice.DLinkedNode)(nil)}

	node.next = node     // nil的变量不可以访问
	node = &DLinkedNode{ // nil的变量可以赋值
		key:   1,
		value: 232,
	}
	fmt.Printf("%#v\n", node)
}
