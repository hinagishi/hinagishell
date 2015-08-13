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
