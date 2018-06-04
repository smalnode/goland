package chan_test

import (
	"testing"
	"time"
)

func And(channels ...<-chan time.Time) <-chan time.Time {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	done := make(chan time.Time)
	go func() {
		defer close(done)
		subdone := And(channels[1:]...)
		for c := 0; c < 2; c++ {
			select {
			case <-channels[0]:
			case <-subdone:
			}
		}
	}()

	return done
}

func TestAnd(t *testing.T) {
	c := And(time.After(100*time.Millisecond),
		time.After(200*time.Millisecond),
		time.After(300*time.Millisecond),
		time.After(300*time.Millisecond),
		time.After(900*time.Millisecond))
	<-c
}
