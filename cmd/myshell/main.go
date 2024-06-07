package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "$ ")

	bufio.NewReader(os.Stdin).ReadString('\n')
}