package radix_tree

import (
    "fmt"
)

func (rt *RadixTree) Insert(key, val string) {
    if(rt == nil || key == "" || val == "") {
        return 
    }
    rt.insertToChild(rt.Root, key, val)
}

func (rt *RadixTree) insertToNode(node *radix_node, key, val string) {
    if node.isRootNode(){
        rt.insertToChild(node, key, val)
    }
    prefixIndex := calcPrefix(key, node.Key)
    // 再次进行验证是否开始字符一致
    if prefixIndex == 0 {
        return 
    }
    fmt.Println("prefix index: ", prefixIndex, node.Key)
    // 需要分叉时， 将分叉点值设为空
    if(prefixIndex != len(node.Key) && prefixIndex != len(key)) {
        node.Val = ""
    }
    if prefixIndex < len(node.Key) {
        oldKey := node.Key
        oldVal := node.Val
        node.Key = oldKey[:prefixIndex]
        if len(node.Childs) == 0 {
            rt.insertToChild(node, oldKey[prefixIndex:], oldVal)
        } else {
            newNode := &radix_node{
                Key: oldKey[prefixIndex:],
                Val: oldVal,
                Childs: node.Childs,
            }
            node.Childs = []*radix_node{newNode}
        }
    }
    if prefixIndex < len(key) {
        rt.insertToChild(node, key[prefixIndex:], val)
    } else {
        // 前缀相等时更改相对应的值
        node.Val = val
    }
}

func (rt *RadixTree) insertToChild(node *radix_node, key, val string) {
    if(len(node.Childs) == 0) {
        newNode := radix_node{
            Key: key,
            Val: val,
        }
        node.Childs = append(node.Childs, &newNode)
        return 
    }
    for i := range node.Childs {
        if isPartPrefix(node.Childs[i].Key, key) {
            rt.insertToNode(node.Childs[i], key, val)
            return
        }
    }
    node.Childs = append(node.Childs, &radix_node{Key: key, Val: val})
}
