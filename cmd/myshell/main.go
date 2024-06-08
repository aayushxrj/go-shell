package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		case "type":
			if len(args) > 1 {
				switch args[1] {
				case "echo", "pwd", "cd","type", "exit":
					fmt.Fprint(os.Stdout,  args[1]+" is a shell builtin\n")
				default:
						pathEnv := os.Getenv("PATH")
						paths := strings.Split(pathEnv, ":")
						found := false
						for _, path := range paths {
							fullPath := filepath.Join(path, args[1])
							_, err := os.Stat(fullPath)
							if err == nil {
								fmt.Fprint(os.Stdout, args[1]+" is "+fullPath+"\n")
								found = true
								break
							}
						}
						if !found {
							fmt.Fprint(os.Stdout, args[1]+": not found\n")
						}
				}
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprint(os.Stdout, err)
			}
			fmt.Fprint(os.Stdout, dir+"\n")
		case "cd":
			if len(args) > 1 {
				if args[1] == "~" {
					err := os.Chdir(os.Getenv("HOME"))
					if err != nil {
						fmt.Fprintln(os.Stdout, err)
					}
				} else {
					err := os.Chdir(args[1])
					if err != nil {
						fmt.Fprintln(os.Stdout, args[1]+": No such file or directory")
					}
				}
			} else {
				err := os.Chdir(os.Getenv("HOME"))
				if err != nil {
					fmt.Fprintln(os.Stdout, err)
				}
			}
		case "exit":
			if len(args) > 1 && args[1] == "0" {
				os.Exit(0)
			}
		default:
			// fmt.Fprint(os.Stdout, input+": command not found\n")
			command := exec.Command(args[0], args[1:]...)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr

			err := command.Run()
			if err != nil {
				fmt.Fprint(os.Stdout, args[0]+": command not found\n")
			}
		}
	}
}