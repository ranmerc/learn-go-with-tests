package spy_countdown_operations

const (
	SleepOperation = "sleep"
	WriteOperation = "write"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, SleepOperation)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, WriteOperation)

	return
}
