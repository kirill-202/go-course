package average

import (
	"fmt"
	"os"
)


func Average(){
	var sum float64
	var n  int

	for {
		var val float64

		if _, err := fmt.Fscanln(os.Stdin, &val); err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Printf("The average is %.2f\n", sum/float64(n))
}