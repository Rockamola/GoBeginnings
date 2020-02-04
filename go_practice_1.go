//go_practice_1 prints command-line arguments, as well as index element and its value, but takes in explicit variable to do so
package main

import (
	//format package for compiling into source files
	"fmt"
	//access to operation system
	"os"
	"time"
)

func main() {
	//explicit declaration for variable
	var s = ""
	//initializing time measurement
	start := time.Now()
	//declared both index and value to return both
	for i, arg := range os.Args[1:] {
		s := s + arg

		fmt.Println("item", i, "is", s)
	}
	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}
