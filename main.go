package main

import (
	"./utils"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type input []string

func (c input) parse() int {
	if c[0] == "pwd" {
		utils.Pwd()
	} else if c[0] == "cd" {
		var path string
		if len(c) < 2 {
			path = os.Getenv("HOME")
		} else {
			path = c[1]
		}
		utils.Cd(path)
	} else if c[0] == "exit" {
		os.Exit(0)
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
		var params input
		params = strings.Split(scanner.Text(), " ")
		params.parse()
	}
}
