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

	var textBuff []byte

	for {

		buf := make([]byte, 1)
		_, err :=	os.Stdin.Read(buf)

		// backspace
		if buf[0] == 127 {
			textBuff = textBuff[:len(textBuff) - 1]
			file.Truncate(0) // size 0b to rewrite
			file.Seek(0,0) // start of the file
			file.Write(textBuff)
			t.Write([]byte("\b \b")) // clean terminal

		} else if len(textBuff) >= 0 {
			textBuff = append(textBuff, buf[0])

			file.Write([]byte{textBuff[len(textBuff) - 1]})
			t.Write([]byte{textBuff[len(textBuff) - 1]})
		}

		if err != nil {
			panic(err)
		}

		if buf[0] == 'q' {
			break
		}

	}

}
