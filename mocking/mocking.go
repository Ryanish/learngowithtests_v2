package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is an implementation of Sleeper with a predefined delay.
type DefaultSleeper struct{}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Sleep will pause execution for the defined Duration.
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
