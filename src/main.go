package main

import (
	"fmt"
	_ "goms/setup"
	"goms/routers"
)

func main() {
	fmt.Println("Begin...")
	r := routers.InitRouters()
	r.Run(":8090")
}