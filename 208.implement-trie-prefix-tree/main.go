package main

import "fmt"

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple")) // 返回 True
	//fmt.Println(trie.Search("app"))     // 返回 False
	//fmt.Println(trie.StartsWith("app")) // 返回 True
	//trie.Insert("app")
	//fmt.Println(trie.Search("app"))
}

type Trie struct {
	children [26]*Trie // 数组
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t // 为什么不直接用t?
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil { // 子节点不存在。创建一个新的子节点，记录在 children 数组的对应位置上，然后沿着指针移动到子节点，继续搜索下一个字符。
			node.children[ch] = &Trie{} // 索引为 ch 的数据赋值为空的节点
		}
		node = node.children[ch] // 移动 node?  子节点存在。沿着指针移动到子节点，继续处理下一个字符。
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
