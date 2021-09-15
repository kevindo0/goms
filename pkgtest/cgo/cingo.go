package cgo

/*
#include <stdio.h>

void pri() {
    printf("hey\n");
}

int add(int a, int b) {
    return a + b;
}
*/
import "C"

import (
    "fmt"
)

func Output1() {
    fmt.Println(C.add(3, 4))
    C.pri()
}

// 7
// hey
