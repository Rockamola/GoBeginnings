//go_practice_1 prints command-line arguments, as well as the command used to invoke file
package main

import (
	//format package for compiling into source files
	"fmt"
	//access to operation system
	"os"
)

func main() {
	//loop over arguments by index and value
	for i, arg := range os.Args {
		//print both the index element and its value
		fmt.Println("item", i, "is", arg)
	}
}
