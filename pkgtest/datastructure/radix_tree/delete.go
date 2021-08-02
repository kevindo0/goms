package radix_tree

import (
    "fmt"
)

func (tr *RadixTree) Delete(key string) bool {
    if key == "" {
        return false
    }
    if tr == nil {
        return false
    }
    return deleteFromChild(tr.Root, key)
}

func deleteFromNode(parent *radix_node, index int, key string) bool {
    node := parent.Childs[index]
    prefixIndex := calcPrefix(key, node.Key)
    // fmt.Println("prefix index:", key, node.Key, prefixIndex)
    newKey := key[prefixIndex:]
    if prefixIndex < len(node.Key) || len(newKey) > 0 {
        fmt.Println("new key:", newKey)
        return deleteFromChild(node, newKey)
    } else {
        if len(node.Childs) == 0 {
            parent.Childs = append(parent.Childs[:index], parent.Childs[index+1:]...)
        } else {
            node.Val = ""
        }
        return true
    }
    return false
}

func deleteFromChild(node *radix_node, key string) bool {
    if len(node.Childs) == 0 {
        return false
    }
    for i := range node.Childs {
        if isPartPrefix(node.Childs[i].Key, key) {
            // fmt.Println("key:", key, i, node.Childs[i].Key)
            return deleteFromNode(node, i, key)
        }
    }
    return false
}

// import (
//     "fmt"
//     radix "pkgte/datastructure/radix_tree"
// )

// func main() {
//     root := radix.NewRadixTree()
//     root.Insert("tester", "1")
//     root.Insert("slow", "2")
//     root.Insert("water", "3")
//     root.Insert("slower", "4")
//     root.Insert("word5", "5")
//     root.Insert("test", "6")
//     fmt.Println(root)
//     fmt.Println(root.Delete("slow"))
//     fmt.Println(root.Delete("slower"))
//     fmt.Println(root.Delete("test"))
//     fmt.Println(root)
// }
