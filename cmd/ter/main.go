package main

import (
	"os"

	"golang.org/x/term"
	"term/internal/ter"

)

func main() {

	argsFile := os.Args[1]

	f, err := os.OpenFile(argsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	state, err := ter.EnableRawMode()
	defer ter.DisableRawMode(state)

	ter.ClearTerminal()

  t := term.NewTerminal(os.Stdout, "")

	ter.Start(t, f)

}
