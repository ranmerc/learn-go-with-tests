package countdown

import (
	"9_mocking/spy_countdown_operations"
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &spy_countdown_operations.SpyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &spy_countdown_operations.SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			spy_countdown_operations.WriteOperation,
			spy_countdown_operations.SleepOperation,
			spy_countdown_operations.WriteOperation,
			spy_countdown_operations.SleepOperation,
			spy_countdown_operations.WriteOperation,
			spy_countdown_operations.SleepOperation,
			spy_countdown_operations.WriteOperation,
		}

		if !reflect.DeepEqual(spySleepPrinter.Calls, want) {
			t.Errorf("wanted calls %v, got calls %v", want, spySleepPrinter.Calls)
		}
	})
}
