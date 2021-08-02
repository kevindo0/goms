package radix_tree

import (
    "fmt"
    "bytes"
    "encoding/json"
)

// 基数树，也称压缩前缀树（例：https://github.com/ZBIGBEAR/radix_tree）

const (
    RootNodeKey = "######"
    RootNodeVal = "######"
)

type radix_node struct {
    Key string                  `json:"Key"`
    Val string                  `json:"Val"`
    Childs []*radix_node        `json:"Childs"`
}

type RadixTree struct {
    Root *radix_node  `json:"Root"`
}

func NewRadixTree() *RadixTree {
    return &RadixTree{
        Root: &radix_node{
            Key: RootNodeKey,
            Val: RootNodeVal,
        },
    }
}

func (rt *radix_node) isRootNode() bool {
    return rt != nil && rt.Val == RootNodeVal && rt.Key == RootNodeKey
}

// 转成json
func (rt *RadixTree) String() string {
    b, err := json.Marshal(rt)
    if err != nil {
        fmt.Println("#error: ", err)
        return ""
    }
    var out bytes.Buffer
    err = json.Indent(&out, b, "", "    ")
    return out.String()
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
// }