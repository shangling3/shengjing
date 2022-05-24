// Echo2 prints its command-like arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seq string
	for _, a := range os.Args[1:] {
		s += a + seq
		seq = " "
	}
	fmt.Println(s)
}
