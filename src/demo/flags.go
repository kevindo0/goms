package main

import (
	"fmt"
	"flag"
	"strconv"
)

const (
	APP_VERSION = "0.1"
)

// 只要实现 flag.Value interface:
//  type Value interface {
//      String() string
//      Set(string) error        
//  }

// 以 flag.Var(&flagvar, name, usage) 方式定义此 flag

type percentage float32
func (p *percentage) Set(s string) error {
    v, err := strconv.ParseFloat(s, 32)
    *p = percentage(v)
    return err
}
func (p *percentage) String() string {
    return fmt.Sprintf("%f", *p)
}

func main() {
	fmt.Println(APP_VERSION)
	namePtr := flag.String("name", "hl", "user`s name")
	agePtr := flag.Uint("age", 22, "user`s age")

	var email string
	flag.StringVar(&email, "email", "2572207199@qq.com", "user`s email")

	var pop percentage
  flag.Var(&pop, "pop", "popularity")

	flag.Parse()
	others := flag.Args()

	fmt.Println("name:", *namePtr)
	fmt.Println("age:", *agePtr)
	fmt.Println("email:", email)
	fmt.Println("pop:", pop)
	fmt.Println(others)
}