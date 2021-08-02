package datastructure

import (
	"strings"
)

// 字典树

/*
给定一定数量的单词，查询给定字符串出现的单词的次数和以此字符串开头的所有单词的数量
构建字典树，也称前缀树
count表示以当前单词结尾的单词数量。
prefix表示以该处节点之前的字符串为前缀的单词数量。
*/
type TrieNode struct {
	count  int
	prefix int
	data   []TrieNode
}

func NewTrieNode() TrieNode {
	data := make([]TrieNode, 26)
	return TrieNode{0, 0, data}
}

// 在字典树中插入数据
func (td *TrieNode) Insert(str string) {
	if len(str) == 0 {
		return
	}
	str = strings.ToLower(str)
	node := td
	bytes := []byte(str)
	for _, val := range bytes {
		index := val - 97
		if len(node.data[index].data) == 0 {
			node.data[index] = NewTrieNode()
		}
		node = &node.data[index]
		node.prefix += 1
	}
	node.count += 1
}

func (td *TrieNode) SearchCount(str string) int {
	if len(str) == 0 {
		return 0
	}
	str = strings.ToLower(str)
	node := td
	bytes := []byte(str)
	for _, val := range bytes {
		index := val - 97
		if len(node.data[index].data) == 0 {
			return 0
		}
		node = &node.data[index]
	}
	return node.count
}

func (td *TrieNode) SearchPrefix(str string) int {
	if len(str) == 0 {
		return 0
	}
	str = strings.ToLower(str)
	node := td
	bytes := []byte(str)
	for _, val := range bytes {
		index := val - 97
		if len(node.data[index].data) == 0 {
			return 0
		}
		node = &node.data[index]
	}
	return node.prefix
}

// import (
//     "fmt"
//     dt "pkgte/datastructure"
// )
// 
// func main() {
// 	root := dt.NewTrieNode()
// 	root.Insert("hello")
// 	root.Insert("hello")
// 	root.Insert("hello")
// 	root.Insert("hello")
// 	root.Insert("helloWorld")
// 	fmt.Println(root.SearchCount("hello"))       // 4
// 	fmt.Println(root.SearchPrefix("hello"))      // 5
// 	fmt.Println(root.SearchPrefix("helloworld")) // 1
// }
