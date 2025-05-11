package main

import (
	"os"

	"golang.org/x/term"
)

/*

0: stdin (entrada estándar) — el flujo de datos para la entrada, generalmente el teclado.

1: stdout (salida estándar) — el flujo de datos para la salida, generalmente la terminal.

2: stderr (error estándar) — para mensajes de error, también generalmente a la terminal.


*/

func main() {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)


	t := term.NewTerminal(os.Stdout, "")

	for {

		_, err := t.ReadLine()

		if err != nil {
			panic(err)
		}

	}


}
