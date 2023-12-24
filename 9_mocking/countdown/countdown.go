package countdown

import (
	"9_mocking/sleeper"
	"fmt"
	"io"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer, s sleeper.Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		s.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
