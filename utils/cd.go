package utils

import (
	"fmt"
	"os"
)

func Cd(path string) {
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
}
