package ter

import (
	"os"

	"golang.org/x/term"
)

func Start(t *term.Terminal, f *os.File) {
	var textBuff []byte

	for {

		buf := make([]byte, 1)
		_, err :=	os.Stdin.Read(buf)

		// backspace
		if buf[0] == 127 {
			textBuff = textBuff[:len(textBuff) - 1]
			f.Truncate(0) // size 0b to rewrite
			f.Seek(0,0) // start of the file
			f.Write(textBuff)
			t.Write([]byte("\b \b")) // clean terminal

		} else if len(textBuff) >= 0 {
			textBuff = append(textBuff, buf[0])

			f.Write([]byte{textBuff[len(textBuff) - 1]})
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

