package radix_tree

import (
    "fmt"
    "bytes"
    "encoding/json"
)

const (
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
            Val: RootNodeVal,
        },
    }
}

func (rt *radix_node) isRootNode() bool {
    return rt != nil && rt.Val == RootNodeVal
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
//     root.Insert("hello", "1")
//     root.Insert("test", "2")
//     root.Insert("tep", "3")
//     fmt.Println(root)
// }