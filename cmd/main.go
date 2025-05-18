package main

import (
	"os"
	"term/internal/ter"

	"golang.org/x/term"
)

func main() {

	argsFile := os.Args[1]

	state, err := ter.EnableRawMode()
	if err != nil {
		panic(err)
	}

	defer ter.DisableRawMode(state)

	ter.ClearTerminal()

  t := term.NewTerminal(os.Stdout, "")

	ter.Start(t, argsFile)

}
