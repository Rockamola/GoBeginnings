//go_dup_check takes multiple user inputs with varying returned results:
//call a file to return the most repeated word, along with its count, or
//create a file via terminal, with results returned upon save
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//empty map
	counts := make(map[string]int)
	//reader for user input
	reader := bufio.NewReader(os.Stdin)
	//holder for user input
	var userInput string
	fmt.Println("Would you like to read from a file? y or n?")
	userInput, _ := reader.ReadString("/n")

}
