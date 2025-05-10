package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/x/term"
)

func main() {
	if term.IsTerminal(os.Stdin.Fd()) {
		fmt.Println("si")
	} else {
		fmt.Println("no")
	}
}
