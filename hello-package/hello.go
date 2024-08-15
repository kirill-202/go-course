package hello

import (
	"fmt"
	"os"
	"strings"
)

func Hello() {
	if len(os.Args) > 1 {
		fmt.Printf("Hello, %s\n!", os.Args[1])
	} else {
		fmt.Println("Hello, world!")
	}
}

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	} 
	return fmt.Sprintf("Hello, "+ strings.Join(names, ", ") + "!")
}