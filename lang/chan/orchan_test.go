package chan_test

import (
	"testing"
	"time"
)

// Or combine list of channels into one channel with or-caluse
func Or(channels ...<-chan time.Time) <-chan time.Time {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	done := make(chan time.Time)
	go func() {
		defer close(done)
		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-Or(append(channels[2:], done)...):
			}
		}
	}()

	return done
}

func TestNil(t *testing.T) {
	c := Or()
	if c != nil {
		t.Error("c is not nil ")
	}
}

func TestOne(t *testing.T) {
	c := Or(time.After(1 * time.Millisecond))
	<-c
}

func TestTwo(t *testing.T) {
	c := Or(time.After(500*time.Millisecond),
		time.After(100*time.Millisecond))
	<-c
}

func TestN(t *testing.T) {
	c := Or(time.After(500*time.Millisecond),
		time.After(300*time.Millisecond),
		time.After(600*time.Millisecond),
		time.After(400*time.Millisecond),
		time.After(900*time.Millisecond))
	<-c
}
