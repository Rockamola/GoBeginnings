//go_practice_1 prints command-line arguments
package main

import (
	//format package for compiling into source files
	"fmt"
	//access to operation system
	"os"
)

func main() {
	//explicit declaration for variables
	s, sep := "", ""
	//_ is used to discard loops index, storing only the element value in arg
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
