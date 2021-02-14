package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Println(cmdString)
	}
}
