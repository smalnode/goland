package lang

import (
	"testing"
	"time"
)

func TestSelectChan(t *testing.T) {
	t5 := time.After(5 * time.Millisecond)
	t3 := time.After(3 * time.Millisecond)
	dl := time.After(10 * time.Millisecond)

outfor:
	for {
		select {
		case <-t5:
			t.Log("t5 return ")
		case <-t3:
			t.Log("t3 return ")
		case <-dl:
			t.Log("times up ")
			break outfor
		}
	}
}

func TestSelectSend(t *testing.T) {
	c := make(chan int, 3)
	count := 0
	// if serval case are all avaiable, then they are execute in persudo-random order
	for {
		select {
		case c <- 0:
		case c <- 1:
		case r := <-c:
			count++
			t.Log(r)
		}
		if count > 20 {
			close(c)
			t.Log("print rest.... ")
			for r := range c {
				t.Log(r)
			}
			break
		}
	}
}
