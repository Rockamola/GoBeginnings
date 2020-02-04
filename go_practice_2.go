//go_practice_2 prints command-line arguments, as well as the command used to invoke file, using string package
package main

import (
	//format package for compiling into source files
	"fmt"
	//imple functions for string manipulation
	"strings"
	//access to operation system
	"os"
	//time monitoring
	"time"
)

func main() {
	//initializing time measurement
	start := time.Now()
	//joining command-line arguments via strings package, passed into empty variable at end
	fmt.Println(strings.Join(os.Args[1:], " "))
	//measuring operation runtime
	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}
