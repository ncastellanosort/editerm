package ter

import (
	"log"
	"os"

	"golang.org/x/term"
)

func ClearTerminal() {
	os.Stdout.Write([]byte("\033[2J\033[H"))
}

func EnableRawMode() (*term.State, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("err making terminal raw %v", err)
	}

	return oldState, nil
}

func DisableRawMode(state *term.State) {
	term.Restore(int(os.Stdin.Fd()), state)
}

