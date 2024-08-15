package main

import (
	"fmt"
	"hello"
	"os"
)

func main(){
	hello := hello.Say(os.Args[1:])
	fmt.Println(hello)
}