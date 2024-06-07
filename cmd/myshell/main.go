package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		input = strings.TrimSpace(input) 

		args := strings.Split(input, " ") 

		switch args[0] {
		case "echo":
			if len(args) > 1 {
				fmt.Fprintln(os.Stdout, strings.Join(args[1:], " "))
			}
		case "exit":
			if len(args) > 1 && args[1] == "0" {
				os.Exit(0)
			}
		default:
			fmt.Fprint(os.Stdout, input+": command not found\n")
		}
	}
}