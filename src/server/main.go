package main

import (
  "fmt"
  "goms/config"
  "goms/ktest"
  "github.com/go-yaml/yaml"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`
type T struct {
    A string
    B struct {
        RenamedC int   `yaml:"c"`
        D        []int `yaml:",flow"`
    }
}

func main() {
    fmt.Println("a")
    fmt.Println(config.A)
    ktest.Ptest()
    t := T{}
    
    err := yaml.Unmarshal([]byte(data), &t)
    if err != nil {
        fmt.Printf("error: %v", err)
    }
    fmt.Printf("--- t:\n%v\n\n", t)
}
