package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

/*

0: stdin (entrada estándar) — el flujo de datos para la entrada, generalmente el teclado.

1: stdout (salida estándar) — el flujo de datos para la salida, generalmente la terminal.

2: stderr (error estándar) — para mensajes de error, también generalmente a la terminal.


*/

func main() {

	file, err := os.Create("created.txt")
	if err != nil {
		panic(err)
	}

	defer os.Remove("created.txt")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	os.Stdout.Write([]byte("\033[2J\033[H"))

  t := term.NewTerminal(os.Stdout, "")

	for {

		// guardar cada letra presionada ASCII
		buf := make([]byte, 1)

		_, err :=	os.Stdin.Read(buf)

		str := fmt.Sprintf("ascii (%c)\n", buf[0])

		file.Write(buf)
		t.Write([]byte(str))

		// backspace
		if buf[0] == 127 {
			buf = buf[:len(buf) - 1]
		}

		if err != nil {
			panic(err)
		}

		if buf[0] == 'q' {
			break
		}

	}

}
