package sleeper

import "time"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func NewConfigurableSleeper(duration time.Duration, sleep func(time.Duration)) *ConfigurableSleeper {
	return &ConfigurableSleeper{duration, sleep}
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
