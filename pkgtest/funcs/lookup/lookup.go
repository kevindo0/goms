package main

import (
	"fmt"
	"strconv"
)

func Lookup(tag string, key string) (value string, ok bool) {
	for tag != "" {
		// Skip leading space.
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		fmt.Println("00 i:", i)

		tag = tag[i:]
		if tag == "" {
			break
		}
		i = 0
		fmt.Println("11 i:", i, ' ', ':', ';', 0x7f)
		for i < len(tag) && tag[i] > ' ' && tag[i] != ':' && tag[i] != '"' && tag[i] != 0x7f {
			i++
		}
		fmt.Println("22 i:", i, string(tag[i]))
		if i == 0 || i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]
		fmt.Println("name:", name, tag)

		// Scan quoted string to find value.
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		if key == name {
			value, err := strconv.Unquote(qvalue)
			if err != nil {
				break
			}
			return value, true
		}
	}
	return "", false
}

func main() {
	tag := `  json:"good"`
	a, err := Lookup(tag, "json")
	fmt.Println(a, err)
}
