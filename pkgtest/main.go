package main

import (
    "fmt"
    "encoding/hex"
)

const hextable = "0123456789abcdef"

func Encode(b []byte) []byte {
    enc := make([]byte, hex.EncodedLen(len(b)))
    hex.Encode(enc, b)
    return enc
}

func main() {
   a := fmt.Sprintf("#{hextable}")
   fmt.Println(a)
}
