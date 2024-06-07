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

		input = input[:len(input)-1]

		switch input {
		case "exit 0":
			os.Exit(0)
		default:
			fmt.Fprint(os.Stdout, input+": command not found\n")
		}
	}
}