package test_interface


import (
	"fmt"
	"io"
	"os"
)
type ByteCounter int

func (bc *ByteCounter) Write(b []byte) (int, error) {
	length := len(b)
	
	*bc += ByteCounter(length)
	return length, nil
}
func TestInterface() {

	var bc ByteCounter
	f1, _ := os.Open("interfaces-package/a.txt")
	f2 := &bc

	n, _ := io.Copy(f2, f1)

	fmt.Println("copied", n, "bytes")
	fmt.Println(bc)
}