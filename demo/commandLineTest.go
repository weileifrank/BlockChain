package main

import (
	"fmt"
	"os"
)

func main() {
	length:=len(os.Args)
	fmt.Printf("cmd len:%d\n",length)
	for i, cmd := range os.Args {
		fmt.Printf("arg[%d] : %s\n", i, cmd)
	}
}
