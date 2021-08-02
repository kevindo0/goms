package main

import (
    "fmt"
    radix "pkgte/datastructure/radix_tree"
)

func main() {
    root := radix.NewRadixTree()
    root.Insert("tester", "1")
    root.Insert("slow", "2")
    root.Insert("water", "3")
    root.Insert("slower", "4")
    root.Insert("word5", "5")
    root.Insert("test", "6")
    fmt.Println(root)
    fmt.Println(root.Delete("slow"))
    fmt.Println(root.Delete("slower"))
    fmt.Println(root.Delete("test"))
    fmt.Println(root)
}