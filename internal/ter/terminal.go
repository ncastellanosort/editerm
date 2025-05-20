package ter

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

type UserTerm struct {
	sysTerm *term.Terminal
	state *term.State
}

func NewUserTerm() *UserTerm {
  s := term.NewTerminal(os.Stdout, "")
	return &UserTerm{
		sysTerm: s,
		state: nil,
	}
}

func (u *UserTerm) WriteText(text []byte) error {
	_, err := u.sysTerm.Write(text)
	if err != nil {
		return fmt.Errorf("err cleaning terminal %w", err)
	}
	return nil
}

func (u *UserTerm) clearTerminal() error {
	 _, err := os.Stdout.Write([]byte("\033[2J\033[H"));

	if err != nil {
		return fmt.Errorf("err cleaning terminal %w", err)
	}

	return nil
}

func (u *UserTerm) enableRawMode() error {
	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalf("err making terminal raw %v", err)
	}
	u.state = state

	return nil
}

func (u *UserTerm) disableRawMode() error {
	if err := term.Restore(int(os.Stdin.Fd()), u.state); err != nil {
		log.Fatalf("err deactivating raw mode %v", err)
	}
	return nil
}

