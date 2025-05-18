package ter

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func ClearTerminal() error {
	 _, err := os.Stdout.Write([]byte("\033[2J\033[H"));

	if err != nil {
		return fmt.Errorf("err cleaning terminal %w", err)
	}

	return nil
}

func EnableRawMode() (*term.State, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("err making terminal raw %v", err)
	}

	return oldState, nil
}

func DisableRawMode(state *term.State) error {
	if err := term.Restore(int(os.Stdin.Fd()), state); err != nil {
		log.Fatalf("err deactivating raw mode %v", err)
	}
	return nil
}

