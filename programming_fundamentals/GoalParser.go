package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")

	return command
}

func main() {
	fmt.Println(interpret("G()(al)"))
}
