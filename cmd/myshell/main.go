package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprintf(os.Stdout, "$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(0)
		}

		fmt.Fprint(os.Stdout, input[:len(input)-1]+": command not found\n")
	}
}