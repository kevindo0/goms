package cpp
/*
#cgo LDFLAGS: -L. -lstdc++
#cgo CXXFLAGS: -std=c++14 -I.
#include "keyTool.h"
*/
import "C"
 
func Sign(){
    C.sign()
}

