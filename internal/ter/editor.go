package ter

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Editor struct {
	userTerminal UserTerm
	file *os.File
	buffer []byte
}

func NewEditor(f *os.File) *Editor {
	userTerminal := NewUserTerm()

	var buffer []byte

	return &Editor{
		userTerminal: *userTerminal,
		file: f,
		buffer: buffer,
	}
}

func (e *Editor) Start() {
	var err error
	defer e.file.Close()
	e.userTerminal.StartTerminal()

	e.buffer, err = data(e.file)
	if err != nil {
		log.Fatalf("err fetching data of the file %v", err)
	}

	e.userTerminal.WriteText(e.buffer)

	var prev byte

	for {

		buf := make([]byte, 1)
		_, err :=	os.Stdin.Read(buf)

	  if buf[0] == 58 { // :
			prev = buf[0]
		}

		if buf[0] == 119 && prev == 58 { // w
			err := save(e.file, e.buffer)
			if err != nil {
				log.Fatalf("err saving data in the file %v", err)
			}
			prev = 0
		} 

		// backspace
		if buf[0] == 127 {
			e.buffer = e.buffer[:len(e.buffer) - 1]
			e.userTerminal.WriteText([]byte("\b \b")) 

		} else if len(e.buffer) >= 0 {
			e.buffer = append(e.buffer, buf[0])

			e.userTerminal.WriteText([]byte{e.buffer[len(e.buffer) - 1]})
		}

		if buf[0] == 113 && prev == 58 {
			prev = 0
			break
		}

		if err != nil {
			log.Fatalf("skill issues %v", err)
		}

	}
}

func data(f *os.File) ([]byte, error) {
	return io.ReadAll(f)
}

func save(f *os.File, buf []byte) error {
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("err truncating file %w", err)
	}

	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("err seeking in file %w", err)
	}

	if _, err := f.Write(buf); err != nil {
		return fmt.Errorf("err writing to file %w", err)
	}

	return nil
}

