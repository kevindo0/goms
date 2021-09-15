package cgo

/*
#include "utils/utils.c"
*/
import "C"
import "fmt"

func Output3() {
    fmt.Println(C.sum(9, 8))
}
