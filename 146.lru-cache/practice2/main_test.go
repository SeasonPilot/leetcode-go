package practice2

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
