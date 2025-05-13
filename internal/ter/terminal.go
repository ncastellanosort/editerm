package ter

import (
	"os"

	"golang.org/x/term"
)

func ClearTerminal() {
	os.Stdout.Write([]byte("\033[2J\033[H"))
}

func EnableRawMode() (*term.State, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	return oldState, nil
}

func DisableRawMode(state *term.State) {
	term.Restore(int(os.Stdin.Fd()), state)
}

