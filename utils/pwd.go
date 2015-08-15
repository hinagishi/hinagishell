package utils

import (
	"fmt"
	"os"
)

func Pwd() {
	fmt.Println(os.Getenv("PWD"))
}
