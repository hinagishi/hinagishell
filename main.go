package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type command []string

func (c command) parse() int {
	if c[0] == "pwd" {
		fmt.Println(os.Getenv("PWD"))
	} else if c[0] == "cd" {
		var path string
		if len(c) < 2 {
			path = os.Getenv("HOME")
		} else {
			path = c[1]
		}
		if err := os.Chdir(path); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Setenv("PWD", path)
	} else {
		com := c[0]
		cmd := exec.Command(com)
		cmd.Env = os.Environ()
		cmd.Args = c
		cmd.Dir = "."
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	return 0
}

func main() {
	prompt := "[" + os.Getenv("USER") + "%] "
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			break
		}
		var params command
		params = strings.Split(scanner.Text(), " ")
		params.parse()
	}
}
